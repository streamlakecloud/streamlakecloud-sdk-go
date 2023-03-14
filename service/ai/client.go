package ai

import "github.com/streamlakecloud/streamlakecloud-sdk-go/base"

type AIClient struct {
	*base.OpenAPI
}

func NewAIClient(httpClient base.HTTPClient) *AIClient {
	return &AIClient{
		base.NewClient(httpClient, ServiceInfo, ApiList),
	}
}

func (c *AIClient) ComposeVideo(req ComposeVideoRequest) (*ComposeVideoResponse, error) {
	resp := &ComposeVideoResponse{}
	err := c.OpenAPI.PostForAPIWithRequestResponse("ComposeVideo", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AIClient) ScanImage(req ScanImageRequest) (*ScanImageResponse, error) {
	resp := &ScanImageResponse{}
	err := c.OpenAPI.PostForAPIWithRequestResponse("ScanImage", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
