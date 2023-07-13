package vod

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/streamlakecloud/streamlakecloud-sdk-go/base"
)

func TestExchangeVodUploadTokenInfo(t *testing.T) {
	client := NewVodClient(nil)
	client.ServiceInfo.Host = HOST_ENDPOINT
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST}

	// stage 1: apply upload token
	req := ApplyUploadTokenRequest{
		MediaType: "jpg",
	}
	resp, err := client.ApplyUploadToken(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("%+v", resp.ResponseData)
	}
	token := resp.ResponseData.UploadSignature

	jc := newMediaCloudClient(nil)
	jc.ServiceInfo.Host = resp.ResponseData.UploadEndpoint
	// step 2: apply for vod upload token
	resp2, err2 := jc.ApplyImageUpload(token, "fake.jpeg")
	if err2 != nil {
		t.Fatalf("%e", err2)
	}
	if len(resp2.ImageToken) == 0 {
		t.Fatalf("empty ImageToken")
	}
	if len(resp2.HttpEndPoints) == 0 {
		t.Fatalf("empty HttpEndPoints")
	}
}

func TestFileUploadToken(t *testing.T) {
	client := NewVodClient(nil)
	client.ServiceInfo.Host = HOST_ENDPOINT
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST}

	// stage 1: apply upload token
	req := ApplyUploadTokenRequest{
		MediaType: "jpg",
	}
	resp, err := client.ApplyUploadToken(req)
	if err != nil {
		t.Fatalf("%e", err)
	}
	token := resp.ResponseData.UploadSignature
	t.Logf("going to use upload token: %s", token)

	jc := newMediaCloudClient(nil)
	jc.ServiceInfo.Host = resp.ResponseData.UploadEndpoint
	// step 2: apply for vod upload token
	t.Logf("calling exchangeVodUploadTokenInfo with token %s", token)
	resp2, err2 := jc.ApplyImageUpload(token, "fake.jpeg")
	if err2 != nil {
		t.Fatalf("%e", err2)
	}
	if len(resp2.ImageToken) == 0 {
		t.Fatalf("empty ImageToken")
	}
	if len(resp2.HttpEndPoints) == 0 {
		t.Fatalf("empty HttpEndPoints")
	}
	uploadToken := resp2.UploadToken
	imageToken := resp2.ImageToken
	httpEndPoint := resp2.HttpEndPoints[0]

	t.Logf("UploadToken: %s\nimageToken: %s\nhttpEndPoint: %s", uploadToken, imageToken, httpEndPoint)

	// step 2: verify by upload token
	verifyReq := VerifyUploadTokenRequest{
		VodUploadToken: uploadToken,
	}

	verifyResp, err2 := client.VerifyUploadToken(verifyReq)
	if err2 != nil {
		t.Fatalf("%e", err2)
	}
	mediaId := verifyResp.ResponseData.MediaId
	if len(mediaId) == 0 {
		t.FailNow()
	}
	t.Logf("UploadToken: %s, MediaId: %s", uploadToken, mediaId)
}

func HttpGetDownload(t *testing.T, url, pattern string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("%e", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http %d, error download %s", resp.StatusCode, url)
	}
	out, err2 := os.CreateTemp(os.TempDir(), pattern)
	if err2 != nil {
		t.Fatalf("%e", err)
	}

	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return out.Name(), err
}

func GetExampleImageFile(t *testing.T) (string, error) {
	// link from https://file-examples.com/index.php/sample-images-download/sample-jpg-download/
	url := "https://cdnfile.corp.kuaishou.com/kc/files/a/post-test-materials/Sample-jpg-image-50kb.jpg"
	return HttpGetDownload(t, url, "*.jpg")
}

func GetExampleVideoFile(t *testing.T) (string, error) {
	// link from https://file-examples.com/index.php/sample-video-files/sample-mp4-files/
	url := "https://cdnfile.corp.kuaishou.com/kc/files/a/post-test-materials/big_buck_bunny_720p_5mb.mp4"
	return HttpGetDownload(t, url, "*.mp4")
}

func TestUploadImageFile(t *testing.T) {
	file, err := GetExampleImageFile(t)
	if err != nil {
		t.Fatalf("%e", err)
	}
	client := NewVodClient(nil)
	client.ServiceInfo.Host = HOST_ENDPOINT
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST}

	// stage 1: apply upload token
	req := ApplyUploadTokenRequest{
		MediaType: filepath.Base(file),
	}
	resp, err2 := client.ApplyUploadToken(req)
	if err2 != nil {
		t.Fatalf("%e", err2)
	}
	token := resp.ResponseData.UploadSignature

	jc := newMediaCloudClient(nil)
	jc.ServiceInfo.Host = resp.ResponseData.UploadEndpoint

	// stage 2: open file and call upload
	uploadToken, err3 := jc.UploadImageFileFromPath(token, file)
	if err3 != nil {
		t.Fatalf("%e", err3)
	}

	// stage 3: get media id
	verifyReq := VerifyUploadTokenRequest{
		VodUploadToken: uploadToken,
	}
	verifyResp, err4 := client.VerifyUploadToken(verifyReq)
	if err4 != nil {
		t.Fatalf("%e", err4)
	} else {
		mediaId := verifyResp.ResponseData.MediaId
		if len(mediaId) == 0 {
			t.FailNow()
		}
		t.Logf("UploadToken: %s, MediaId: %s", uploadToken, mediaId)
	}
}

func TestUploadVideoFileWithoutCover(t *testing.T) {
	videoFile, err := GetExampleVideoFile(t)
	if err != nil {
		t.Fatalf("%e", err)
	}
	client := NewVodClient(nil)
	client.ServiceInfo.Host = HOST_ENDPOINT
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST}

	// stage 1: apply upload token
	req := ApplyUploadTokenRequest{
		MediaType: "mp4",
	}
	resp, err2 := client.ApplyUploadToken(req)
	if err2 != nil {
		t.Fatalf("%e", err2)
	}
	token := resp.ResponseData.UploadSignature

	jc := newMediaCloudClient(nil)
	jc.ServiceInfo.Host = resp.ResponseData.UploadEndpoint
	// stage 2: open file and call upload
	videoIO, fileopenerr := os.Open(videoFile)
	if fileopenerr != nil {
		t.Fatalf("%e", fileopenerr)
	}
	uploadToken, err3 := jc.UploadVideoFileWithOptionalCover(token, videoIO, nil, filepath.Base(videoFile), "")
	if err3 != nil {
		t.Fatalf("%e", err3)
	}

	// stage 3: get media id
	verifyReq := VerifyUploadTokenRequest{
		VodUploadToken: uploadToken,
	}
	verifyResp, err4 := client.VerifyUploadToken(verifyReq)
	if err4 != nil {
		t.Fatalf("%e", err4)
	} else {
		mediaId := verifyResp.ResponseData.MediaId
		if len(mediaId) == 0 {
			t.FailNow()
		}
		t.Logf("UploadToken: %s, MediaId: %s", uploadToken, mediaId)
	}
}

func TestUploadVideoFileWithCover(t *testing.T) {
	videoFile, err := GetExampleVideoFile(t)
	if err != nil {
		t.Fatalf("%e", err)
	}
	coverFile, errImageDownload := GetExampleImageFile(t)
	if errImageDownload != nil {
		t.Fatalf("%e", errImageDownload)
	}
	client := NewVodClient(nil)
	client.ServiceInfo.Host = HOST_ENDPOINT
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST}

	// stage 1: apply upload token
	req := ApplyUploadTokenRequest{
		MediaType: "mp4",
	}
	resp, err2 := client.ApplyUploadToken(req)
	if err2 != nil {
		t.Fatalf("%e", err2)
	}
	token := resp.ResponseData.UploadSignature

	jc := newMediaCloudClient(nil)
	jc.ServiceInfo.Host = resp.ResponseData.UploadEndpoint
	// stage 2: open file and call upload
	videoIO, videoOpenError := os.Open(videoFile)
	if videoOpenError != nil {
		t.Fatalf("%e", videoOpenError)
	}
	coverIO, coverOpenError := os.Open(coverFile)
	if coverOpenError != nil {
		t.Fatalf("%e", coverOpenError)
	}
	videoFileName := filepath.Base(videoFile)
	coverFileName := filepath.Base(coverFile)
	uploadToken, err3 := jc.UploadVideoFileWithOptionalCover(token, videoIO, coverIO, videoFileName, coverFileName)
	if err3 != nil {
		t.Fatalf("%e", err3)
	}

	// stage 3: get media id
	verifyReq := VerifyUploadTokenRequest{
		VodUploadToken: uploadToken,
	}
	verifyResp, err4 := client.VerifyUploadToken(verifyReq)
	if err4 != nil {
		t.Fatalf("%e", err4)
	} else {
		mediaId := verifyResp.ResponseData.MediaId
		if len(mediaId) == 0 {
			t.FailNow()
		}
		t.Logf("UploadToken: %s, MediaId: %s", uploadToken, mediaId)
	}
}
