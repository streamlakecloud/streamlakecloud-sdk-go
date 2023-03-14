package vod

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/streamlakecloud/streamlakecloud-sdk-go/base"
)

type vodUploadTokenRequest struct {
	Signature string  `json:"signature"`
	ImageName *string `json:"image_name,omitempty"`
	VideoName *string `json:"video_name,omitempty"`
	CoverName *string `json:"cover_name,omitempty"`
}

type vodUploadTokenResponse struct {
	Result      int32  `json:"result"` // 1 for success
	UploadToken string `json:"upload_token,omitempty"`
	ImageToken  string `json:"image_token,omitempty"`
	VideoToken  string `json:"video_token,omitempty"`
	CoverToken  string `json:"cover_token,omitempty"`
	EndPoints   []struct {
		Protocol string `json:"protocol,omitempty"`
		Host     string `json:"host,omitempty"`
		Port     int32  `json:"port,omitempty"`
	} `json:"endpoint,omitempty"`
	HttpEndPoints []string `json:"http_endpoint,omitempty"`
}

type vodUploadPublishRequest struct {
	UploadToken string `json:"upload_token,omitempty"`
}

type vodUploadPublishResponse struct {
	Result int32 `json:"result"` // 1 for success
}

type vodUploadFragmentRequest struct {
	ContentMd5  string    `json:"content_md5"`
	FragmentId  int64     `json:"fragment_id"`
	UploadToken string    `json:"upload_token"`
	Reader      io.Reader `json:"-"`
}

type vodUploadFragmentResponse struct {
	Result   int32  `json:"result"` // 1 for success
	Checksum string `json:"checksum"`
	Size     uint64 `json:"size"`
}

type vodUploadCompleteRequest struct {
	FragmentCount int64  `json:"fragment_count"`
	UploadToken   string `json:"upload_token"`
}

type vodUploadCompleteResponse struct {
	Result int32 `json:"result"` // 1 for success
}

type uploadType int

const (
	UploadTypeImage uploadType = iota
	UploadTypeVideo
)

const (
	DefaultFragmentSize int64 = 1 << 20 // default fragment size 1M
)

var mediaCloudServiceInfo = base.ServiceInfo{
	Region:      "cn-beijing",
	Host:        "",
	Scheme:      "https",
	Credentials: base.Credentials{},
}

var mediaCloudApiList = map[string]*base.ApiInfo{
	"ApplyImageUpload": {
		Method: http.MethodPost,
		Path:   "/api/upload/apply_image_upload",
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
	},
	"ApplyVideoUpload": {
		Method: http.MethodPost,
		Path:   "/api/upload/apply_video_upload",
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
	},
	"PublishImage": {
		Method: http.MethodPost,
		Path:   "/api/upload/publish_image",
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
	},
	"PublishVideo": {
		Method: http.MethodPost,
		Path:   "/api/upload/publish_video",
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
	},
}

type mediaCloudClient struct {
	*base.OpenAPI
}

func newMediaCloudClient(httpClient base.HTTPClient) *mediaCloudClient {
	return &mediaCloudClient{
		base.NewClient(
			httpClient,
			mediaCloudServiceInfo,
			mediaCloudApiList,
		),
	}
}

var streamLakeUploadServiceInfo = base.ServiceInfo{
	Region:      "cn-beijing",
	Host:        "", // to be set by client per request
	Scheme:      "https",
	Credentials: base.Credentials{},
}

var streamLakeUploadApiList = map[string]*base.ApiInfo{
	"UploadFragment": {
		Method: http.MethodPost,
		Path:   "/api/upload/fragment",
		Header: http.Header{
			"Content-Type": []string{"application/octet-stream"},
		},
	},
	"UploadComplete": {
		Method: http.MethodPost,
		Path:   "/api/upload/complete",
	},
}

type uploadClient struct {
	*base.OpenAPI
}

func newUploadClient(httpClient base.HTTPClient) *uploadClient {
	return &uploadClient{
		base.NewClient(
			httpClient,
			streamLakeUploadServiceInfo,
			streamLakeUploadApiList,
		),
	}
}

func (u *uploadClient) UploadFragment(req vodUploadFragmentRequest) (*vodUploadFragmentResponse, error) {
	resp := &vodUploadFragmentResponse{}
	err := u.PostForAPIWithRequestResponse("UploadFragment", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *uploadClient) UploadComplete(req vodUploadCompleteRequest) (*vodUploadCompleteResponse, error) {
	resp := &vodUploadCompleteResponse{}
	err := u.PostForAPIWithRequestResponse("UploadComplete", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *uploadClient) uploadFragment(uploadToken string, contentMd5 string, fragmentId int64, fileIO io.ReadSeekCloser, fragmentSize int64) error {
	req := vodUploadFragmentRequest{
		UploadToken: uploadToken,
		ContentMd5:  contentMd5,
		FragmentId:  fragmentId,
		Reader:      io.LimitReader(fileIO, fragmentSize),
	}
	resp, err := u.UploadFragment(req)
	if err != nil {
		return err
	}
	if resp.Result != 1 {
		return fmt.Errorf("failed to upload, result != 1")
	}
	return nil
}

func (uc *uploadClient) singleFileFragmentedUploadAndComplete(singleFileIO io.ReadSeekCloser, fragmentSize int64, singleFileToken string) error {
	// seek to correct file pos
	totalSize, err := singleFileIO.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}
	fragmentCount := totalSize / fragmentSize
	if totalSize-fragmentCount*fragmentSize > 0 {
		fragmentCount += 1
	}
	for i := int64(0); i < fragmentCount; i++ {
		currentFragmentSize := fragmentSize
		if currentPos, err2 := singleFileIO.Seek(i*fragmentSize, io.SeekStart); err2 != nil {
			return err2
		} else if totalSize-currentPos < fragmentSize {
			// last fragment size < fragmentSize
			currentFragmentSize = totalSize - currentPos
		}
		// calculate md5 for the next currentFragmentSize bytes
		h := md5.New()
		if written, err3 := io.CopyN(h, singleFileIO, currentFragmentSize); written != currentFragmentSize {
			return fmt.Errorf("failed to calculate md5, internal error: %e", err3)
		}
		md5SumBytes := h.Sum(nil)
		md5SumString := fmt.Sprintf("%x", md5SumBytes)

		if _, err4 := singleFileIO.Seek(i*fragmentSize, io.SeekStart); err4 != nil {
			return err4
		}
		// call upload
		if err5 := uc.uploadFragment(
			singleFileToken,
			md5SumString,
			i,
			singleFileIO,
			currentFragmentSize,
		); err5 != nil {
			return err5
		}
	}
	return uc.uploadComplete(singleFileToken, fragmentCount)
}

func (u *uploadClient) uploadComplete(uploadToken string, fragmentCount int64) error {
	req := vodUploadCompleteRequest{
		UploadToken:   uploadToken,
		FragmentCount: fragmentCount,
	}
	resp, err := u.UploadComplete(req)
	if err != nil {
		return err
	}
	if resp.Result != 1 {
		return fmt.Errorf("failed to complete, result != 1")
	}
	return nil
}

func (jc *mediaCloudClient) ApplyImageUpload(uploadSignature, filename string) (*vodUploadTokenResponse, error) {
	req := vodUploadTokenRequest{
		Signature: uploadSignature,
		ImageName: &filename,
	}
	resp := new(vodUploadTokenResponse)
	if err := jc.PostForAPIWithRequestResponse("ApplyImageUpload", req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (jc *mediaCloudClient) ApplyVideoUpload(uploadSignature, filename, coverFilename string) (*vodUploadTokenResponse, error) {
	req := vodUploadTokenRequest{
		Signature: uploadSignature,
		VideoName: &filename,
	}
	if len(coverFilename) > 0 {
		req.CoverName = &coverFilename
	}
	resp := new(vodUploadTokenResponse)
	if err := jc.PostForAPIWithRequestResponse("ApplyVideoUpload", req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (jc *mediaCloudClient) PublishImage(uploadToken string) (*vodUploadPublishResponse, error) {
	req := vodUploadPublishRequest{
		UploadToken: uploadToken,
	}
	resp := new(vodUploadPublishResponse)
	if err := jc.PostForAPIWithRequestResponse("PublishImage", req, resp); err != nil {
		return nil, err
	}
	if resp.Result != 1 {
		return nil, fmt.Errorf("failed to PublishImage, result != 1")
	}
	return resp, nil
}

func (jc *mediaCloudClient) PublishVideo(uploadToken string) (*vodUploadPublishResponse, error) {
	req := vodUploadPublishRequest{
		UploadToken: uploadToken,
	}
	resp := new(vodUploadPublishResponse)
	if err := jc.PostForAPIWithRequestResponse("PublishVideo", req, resp); err != nil {
		return nil, err
	}
	if resp.Result != 1 {
		return nil, fmt.Errorf("failed to PublishVideo, result != 1")
	}
	return resp, nil
}

func (v *mediaCloudClient) UploadImageFileFromPath(uploadToken string, fullPath string) (string, error) {
	if fi, err := os.Stat(fullPath); err != nil {
		return "", err
	} else if fi.IsDir() {
		return "", fmt.Errorf("not a file: %s", fullPath)
	}
	fileIO, err := os.Open(fullPath)
	if err != nil {
		return "", err
	}
	defer func() {
		fileIO.Close()
	}()
	filename := filepath.Base(fullPath)

	return v.UploadImageFile(uploadToken, fileIO, filename)
}

// UploadImageFile provides native *image* file upload mechanism for VodClient.
// signature can be retrieved by call the VodClient.ApplyUploadToken method.
func (v *mediaCloudClient) UploadImageFile(signature string, imageFileIO io.ReadSeekCloser, filename string) (string, error) {
	resp, err := v.ApplyImageUpload(signature, filename)
	if err != nil {
		return "", err
	}
	vodUploadToken := resp.UploadToken
	imageToken := resp.ImageToken
	httpEndPoint := resp.HttpEndPoints[0]

	uc := newUploadClient(base.DefaultClient)
	uc.ServiceInfo.Host = httpEndPoint
	fragmentSize := DefaultFragmentSize

	err2 := uc.singleFileFragmentedUploadAndComplete(imageFileIO, fragmentSize, imageToken)
	if err2 != nil {
		return vodUploadToken, err2
	}

	_, err3 := v.PublishImage(vodUploadToken)
	if err3 != nil {
		return vodUploadToken, err3
	}
	return vodUploadToken, nil
}

// UploadVideoFileWithOptionalCover provides native *video* and cover image file upload mechanism.
// signature can be retrieved by call the ApplyUploadToken method.
//
// videoFileIO and coverFileIO can be created by os.Open
// if coverFilename is empty, cover is omitted
// The returned VodUploadToken can be used to get MediaId
func (v *mediaCloudClient) UploadVideoFileWithOptionalCover(signature string, videoFileIO, coverFileIO io.ReadSeekCloser, videoFilename, coverFilename string) (string, error) {
	resp, err := v.ApplyVideoUpload(signature, videoFilename, coverFilename)
	if err != nil {
		return "", fmt.Errorf("failed to apply video upload, %e", err)
	}
	vodUploadToken := resp.UploadToken
	videoToken := resp.VideoToken
	coverToken := resp.CoverToken
	httpEndPoint := resp.HttpEndPoints[0]

	uc := newUploadClient(nil)
	uc.ServiceInfo.Host = httpEndPoint
	fragmentSize := DefaultFragmentSize

	err2 := uc.singleFileFragmentedUploadAndComplete(videoFileIO, fragmentSize, videoToken)
	if err2 != nil {
		return vodUploadToken, err2
	}
	if coverFileIO != nil {
		err3 := uc.singleFileFragmentedUploadAndComplete(coverFileIO, fragmentSize, coverToken)
		if err3 != nil {
			return vodUploadToken, err3
		}
	}
	_, err4 := v.PublishVideo(vodUploadToken)
	if err4 != nil {
		return vodUploadToken, err4
	}
	return vodUploadToken, nil
}

func (v *VodClient) UploadImageFile(imageFileIO io.ReadSeekCloser, filename string) (string, error) {
	ext := filepath.Ext(filename)
	if len(ext) == 0 {
		return "", fmt.Errorf("%s has no ext", filename)
	}
	req := ApplyUploadTokenRequest{
		MediaType: ext,
	}
	resp, err := v.ApplyUploadToken(req)
	if err != nil {
		return "", err
	}
	uploadEndpoint := resp.ResponseData.UploadEndpoint
	uploadSignature := resp.ResponseData.UploadSignature
	jc := newMediaCloudClient(nil)
	jc.ServiceInfo.Host = uploadEndpoint
	return jc.UploadImageFile(uploadSignature, imageFileIO, filename)
}

func getFileIOAndBase(fullPath string) (*os.File, string, error) {
	if fi, err := os.Stat(fullPath); err != nil {
		return nil, "", err
	} else if fi.IsDir() {
		return nil, "", fmt.Errorf("not a file: %s", fullPath)
	}
	fileIO, err := os.Open(fullPath)
	if err != nil {
		return nil, "", err
	}
	filename := filepath.Base(fullPath)
	return fileIO, filename, nil
}

func (v *VodClient) UploadImageFileFromPath(fullPath string) (string, error) {
	fileIO, filename, err := getFileIOAndBase(fullPath)
	if err != nil {
		return "", err
	}
	defer func() {
		fileIO.Close()
	}()
	return v.UploadImageFile(fileIO, filename)
}

func (v *VodClient) UploadVideoFileWithOptionalCover(videoFileIO, coverFileIO io.ReadSeekCloser, videoFilename, coverFilename string) (string, error) {
	ext := filepath.Ext(videoFilename)
	if len(ext) == 0 {
		return "", fmt.Errorf("%s has no ext", videoFilename)
	}
	req := ApplyUploadTokenRequest{
		MediaType: ext,
	}
	resp, err := v.ApplyUploadToken(req)
	if err != nil {
		return "", err
	}
	uploadEndpoint := resp.ResponseData.UploadEndpoint
	uploadSignature := resp.ResponseData.UploadSignature
	jc := newMediaCloudClient(nil)
	jc.ServiceInfo.Host = uploadEndpoint
	return jc.UploadVideoFileWithOptionalCover(uploadSignature, videoFileIO, coverFileIO, videoFilename, coverFilename)
}

func (v *VodClient) UploadVideoFileWithOptionalCoverFromPath(videoPath, coverPath string) (string, error) {
	videoFileIO, videoFilename, err := getFileIOAndBase(videoPath)
	if err != nil {
		return "", err
	}

	defer func() {
		videoFileIO.Close()
	}()

	var coverFileIO *os.File
	var coverFilename string

	if len(coverPath) > 0 {
		coverFileIO, coverFilename, err = getFileIOAndBase(coverPath)
		if err != nil {
			return "", err
		}
	}

	defer func() {
		coverFileIO.Close()
	}()

	return v.UploadVideoFileWithOptionalCover(videoFileIO, coverFileIO, videoFilename, coverFilename)
}
