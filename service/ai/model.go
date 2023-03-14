package ai

import "github.com/streamlakecloud/streamlakecloud-sdk-go/base"

type ComposeVideoRequest struct {
	Template    string
	InputURLSet string
	CallbackURL string
	ExtraInfo   string
}

type ComposeVideoResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData interface{} `json:",omitempty"`
}

type ImageSet struct {
	ImageURL    string // ImageURL is the url of the image
	ImageBase64 string
	DataId      string
}

type ScanImageRequest struct {
	ImageSets []ImageSet
	Scenes    string // Scenes is a comma separated string, e.g. "aSceneName,anotherSceneName,YetAnotherSceneName". For now only "porn" is supported
}

type ResultDetail struct {
	Success      bool
	ErrorMessage string
	ErrorCode    string
	Scene        string
	Label        string
	Probability  float32 // 0.0f~1.0f
}

type ScanResult struct {
	TaskId        string
	DataId        string
	ImageURL      string
	ResultDetails []ResultDetail
}

type ScanImageResult struct {
	ScanResults []ScanResult
}

type ScanImageResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData ScanImageResult `json:",omitempty"`
}
