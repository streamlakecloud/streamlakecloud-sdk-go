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

func (v *VodClient) DeleteMedia(req DeleteMediaRequest) (*DeleteMediaResponse, error) {
	resp := &DeleteMediaResponse{}
	err := v.PostForAPIWithRequestResponse("DeleteMedia", req, resp)
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

func (v *VodClient) CreateTranscodeTemplate(req CreateTranscodeTemplateRequest) (*CreateTranscodeTemplateResponse, error) {
	resp := &CreateTranscodeTemplateResponse{}
	err := v.PostForAPIWithRequestResponse("CreateTranscodeTemplate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) UpdateTranscodeTemplate(req UpdateTranscodeTemplateRequest) (*UpdateTranscodeTemplateResponse, error) {
	resp := &UpdateTranscodeTemplateResponse{}
	err := v.PostForAPIWithRequestResponse("UpdateTranscodeTemplate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) DescribeTranscodeTemplate(req DescribeTranscodeTemplateRequest) (*DescribeTranscodeTemplateResponse, error) {
	resp := &DescribeTranscodeTemplateResponse{}
	err := v.PostForAPIWithRequestResponse("DescribeTranscodeTemplate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) ListTranscodeTemplate(req ListTranscodeTemplateRequest) (*ListTranscodeTemplateResponse, error) {
	resp := &ListTranscodeTemplateResponse{}
	err := v.PostForAPIWithRequestResponse("ListTranscodeTemplate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) DeleteTranscodeTemplate(req DeleteTranscodeTemplateRequest) (*DeleteTranscodeTemplateResponse, error) {
	resp := &DeleteTranscodeTemplateResponse{}
	err := v.PostForAPIWithRequestResponse("DeleteTranscodeTemplate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) CreateWatermarkTemplate(req CreateWatermarkTemplateRequest) (*CreateWatermarkTemplateResponse, error) {
	resp := &CreateWatermarkTemplateResponse{}
	err := v.PostForAPIWithRequestResponse("CreateWatermarkTemplate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) UpdateWatermarkTemplate(req UpdateWatermarkTemplateRequest) (*UpdateWatermarkTemplateResponse, error) {
	resp := &UpdateWatermarkTemplateResponse{}
	err := v.PostForAPIWithRequestResponse("UpdateWatermarkTemplate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) DescribeWatermarkTemplate(req DescribeWatermarkTemplateRequest) (*DescribeWatermarkTemplateResponse, error) {
	resp := &DescribeWatermarkTemplateResponse{}
	err := v.PostForAPIWithRequestResponse("DescribeWatermarkTemplate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) ListWatermarkTemplate(req ListWatermarkTemplateRequest) (*ListWatermarkTemplateResponse, error) {
	resp := &ListWatermarkTemplateResponse{}
	err := v.PostForAPIWithRequestResponse("ListWatermarkTemplate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) DeleteWatermarkTemplate(req DeleteWatermarkTemplateRequest) (*DeleteWatermarkTemplateResponse, error) {
	resp := &DeleteWatermarkTemplateResponse{}
	err := v.PostForAPIWithRequestResponse("DeleteWatermarkTemplate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) CreateSnapshotTemplate(req CreateSnapshotTemplateRequest) (*CreateSnapshotTemplateResponse, error) {
	resp := &CreateSnapshotTemplateResponse{}
	err := v.PostForAPIWithRequestResponse("CreateSnapshotTemplate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) UpdateSnapshotTemplate(req UpdateSnapshotTemplateRequest) (*UpdateSnapshotTemplateResponse, error) {
	resp := &UpdateSnapshotTemplateResponse{}
	err := v.PostForAPIWithRequestResponse("UpdateSnapshotTemplate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) DescribeSnapshotTemplate(req DescribeSnapshotTemplateRequest) (*DescribeSnapshotTemplateResponse, error) {
	resp := &DescribeSnapshotTemplateResponse{}
	err := v.PostForAPIWithRequestResponse("DescribeSnapshotTemplate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) ListSnapshotTemplate(req ListSnapshotTemplateRequest) (*ListSnapshotTemplateResponse, error) {
	resp := &ListSnapshotTemplateResponse{}
	err := v.PostForAPIWithRequestResponse("ListSnapshotTemplate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) DeleteSnapshotTemplate(req DeleteSnapshotTemplateRequest) (*DeleteSnapshotTemplateResponse, error) {
	resp := &DeleteSnapshotTemplateResponse{}
	err := v.PostForAPIWithRequestResponse("DeleteSnapshotTemplate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) CreateWorkflowTemplate(req CreateWorkflowTemplateRequest) (*CreateWorkflowTemplateResponse, error) {
	resp := &CreateWorkflowTemplateResponse{}
	err := v.PostForAPIWithRequestResponse("CreateWorkflowTemplate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) UpdateWorkflowTemplate(req UpdateWorkflowTemplateRequest) (*UpdateWorkflowTemplateResponse, error) {
	resp := &UpdateWorkflowTemplateResponse{}
	err := v.PostForAPIWithRequestResponse("UpdateWorkflowTemplate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) ListWorkflowTemplate(req ListWorkflowTemplateRequest) (*ListWorkflowTemplateResponse, error) {
	resp := &ListWorkflowTemplateResponse{}
	err := v.PostForAPIWithRequestResponse("ListWorkflowTemplate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) DeleteWorkflowTemplate(req DeleteWorkflowTemplateRequest) (*DeleteWorkflowTemplateResponse, error) {
	resp := &DeleteWorkflowTemplateResponse{}
	err := v.PostForAPIWithRequestResponse("DeleteWorkflowTemplate", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) DescribeTaskDetail(req DescribeTaskDetailRequest) (*DescribeTaskDetailResponse, error) {
	resp := &DescribeTaskDetailResponse{}
	err := v.PostForAPIWithRequestResponse("DescribeTaskDetail", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) ApplyUploadInfo(req ApplyUploadInfoRequest) (*ApplyUploadInfoResponse, error) {
	resp := &ApplyUploadInfoResponse{}
	err := v.PostForAPIWithRequestResponse("ApplyUploadInfo", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VodClient) CommitUpload(req CommitUploadRequest) (*CommitUploadResponse, error) {
	resp := &CommitUploadResponse{}
	err := v.PostForAPIWithRequestResponse("CommitUpload", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
