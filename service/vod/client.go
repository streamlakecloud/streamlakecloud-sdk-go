package vod

import (
	"github.com/streamlakecloud/streamlakecloud-sdk-go/base"
)

type VodClient struct {
	*base.OpenAPI
}

func NewVodClient(httpClient base.HTTPClient) *VodClient {
	return &VodClient{
		base.NewClient(httpClient, ServiceInfo, ApiList),
	}
}

func NewVodClientV2(httpClient base.HTTPClient, serviceInfo base.ServiceInfo) *VodClient {
	return &VodClient{
		base.NewClient(httpClient, serviceInfo, ApiList),
	}
}

func (v *VodClient) ApplyUploadToken(req ApplyUploadTokenRequest) (*ApplyUploadTokenResponse, error) {
	resp := &ApplyUploadTokenResponse{}
	err := v.PostForAPIWithRequestResponse("ApplyUploadToken", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) VerifyUploadToken(req VerifyUploadTokenRequest) (*VerifyUploadTokenResponse, error) {
	resp := &VerifyUploadTokenResponse{}
	err := v.PostForAPIWithRequestResponse("VerifyUploadToken", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) DescribeFetchJobs(req DescribeFetchJobsRequest) (*DescribeFetchJobsResponse, error) {
	resp := &DescribeFetchJobsResponse{}
	err := v.PostForAPIWithRequestResponse("DescribeFetchJobs", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) FetchStore(req FetchStoreRequest) (*FetchStoreResponse, error) {
	resp := &FetchStoreResponse{}
	err := v.PostForAPIWithRequestResponse("FetchStore", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) FetchUpload(req FetchUploadRequest) (*FetchUploadResponse, error) {
	resp := &FetchUploadResponse{}
	err := v.PostForAPIWithRequestResponse("FetchUpload", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) SubmitWorkflow(req SubmitWorkflowRequest) (*SubmitWorkflowResponse, error) {
	resp := &SubmitWorkflowResponse{}
	err := v.PostForAPIWithRequestResponse("SubmitWorkflow", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) ProcessMedia(req ProcessMediaRequest) (*ProcessMediaResponse, error) {
	resp := &ProcessMediaResponse{}
	err := v.PostForAPIWithRequestResponse("ProcessMedia", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) SubmitMediaProcessJobs(req SubmitMediaProcessJobsRequest) (*SubmitMediaProcessJobsResponse, error) {
	resp := &SubmitMediaProcessJobsResponse{}
	err := v.PostForAPIWithRequestResponse("SubmitMediaProcessJobs", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) DescribeMediaProcessJobs(req DescribeMediaProcessJobsRequest) (*DescribeMediaProcessJobsResponse, error) {
	resp := &DescribeMediaProcessJobsResponse{}
	err := v.PostForAPIWithRequestResponse("DescribeMediaProcessJobs", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) DescribeMediaInfo(req DescribeMediaInfoRequest) (*DescribeMediaInfoResponse, error) {
	resp := &DescribeMediaInfoResponse{}
	err := v.PostForAPIWithRequestResponse("DescribeMediaInfo", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) DescribeAttachedMediaInfo(req DescribeAttachedMediaInfoRequest) (*DescribeAttachedMediaInfoResponse, error) {
	resp := &DescribeAttachedMediaInfoResponse{}
	err := v.PostForAPIWithRequestResponse("DescribeAttachedMediaInfo", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) DescribePlayQualityDataSources(req DescribePlayQualityDataSourcesRequest) (*DescribePlayQualityDataSourcesResponse, error) {
	resp := &DescribePlayQualityDataSourcesResponse{}
	err := v.PostForAPIWithRequestResponse("DescribePlayQualityDataSources", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) DescribePlayQualitySummary(req DescribePlayQualitySummaryRequest) (*DescribePlayQualitySummaryResponse, error) {
	resp := &DescribePlayQualitySummaryResponse{}
	err := v.PostForAPIWithRequestResponse("DescribePlayQualitySummary", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) DescribePlayQualityDetail(req DescribePlayQualityDetailRequest) (*DescribePlayQualityDetailResponse, error) {
	resp := &DescribePlayQualityDetailResponse{}
	err := v.PostForAPIWithRequestResponse("DescribePlayQualityDetail", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) DetectMedia(req DetectMediaRequest) (*DetectMediaResponse, error) {
	resp := &DetectMediaResponse{}
	err := v.PostForAPIWithRequestResponse("DetectMedia", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
