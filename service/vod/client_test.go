package vod

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/streamlakecloud/streamlakecloud-sdk-go/base"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

var (
	HOST_ENDPOINT   = os.Getenv("STREAMLAKE_VOD_GO_SDK_ENDPOINT")
	ACCESS_KEY_TEST = os.Getenv("STREAMLAKE_VOD_GO_SDK_ACCESS_KEY2")
	SECRET_KEY_TEST = os.Getenv("STREAMLAKE_VOD_GO_SDK_SECRET_KEY")
	MEDIA_ID        = "aa99ded2648e078d"
)

func TestApplyUploadToken(t *testing.T) {
	client := NewVodClient(nil)
	client.ServiceInfo.Host = HOST_ENDPOINT
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST}
	client.ServiceInfo.ProductName = "vod"
	client.ServiceInfo.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req := ApplyUploadTokenRequest{
		MediaType: "jpg",
	}
	resp, err := client.ApplyUploadToken(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestFetchUpload(t *testing.T) {
	client := NewVodClient(nil)
	client.ServiceInfo.Host = HOST_ENDPOINT
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST}
	req := FetchUploadRequest{
		URLSets: []FetchUploadURLSet{
			{
				SourceURL:    "https://static.streamlake.com/kos/nlav11935/streamlake-website/static/home/main.mp4",
				CallbackArgs: "test",
				CallbackURL:  "www.streamlake.com",
			},
		},
	}
	resp, err := client.FetchUpload(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestDescribeFetchJobs(t *testing.T) {
	client := NewVodClient(nil)
	client.ServiceInfo.Host = HOST_ENDPOINT
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST}
	req := DescribeFetchJobsRequest{
		JobIds: []string{"93dfdb5b-e643-40d5-9705-f22fca835976-20220615-0"},
	}
	resp, err := client.DescribeFetchJobs(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestFetchStore(t *testing.T) {
	client := NewVodClient(nil)
	client.ServiceInfo.Host = HOST_ENDPOINT
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST}
	req := FetchStoreRequest{
		URLSets: []FetchStoreURLSet{
			{
				SourceURL:    "",
				StoreKey:     "big_buck_bunny_720p_5mb.mp4",
				CallbackArgs: "test",
			},
		},
	}
	resp, err := client.FetchStore(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

// 获取媒资信息列表样例
func TestDescribeListMediaInfo(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: "ACCESS_KEY_TEST", SecretAccessKey: "SECRET_KEY_TEST"},
	}
	client := NewVodClientV2(nil, serviceInfo)
	req := ListMediaInfoRequest{}
	resp, err := client.ListMediaInfo(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

// 获取媒资信息样例
func TestDescribeMediaInfo(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)
	req := DescribeMediaInfoRequest{
		MediaId: MEDIA_ID,
		Filters: []string{},
	}
	resp, err := client.DescribeMediaInfo(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

// 删除媒资样例
func TestDeleteMedia(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)
	req := DeleteMediaRequest{
		MediaId: MEDIA_ID,
		DeleteItems: []MediaDeleteItem{
			{
				Type:       "TranscodeFilesByTemplate",
				TemplateId: "540P-MP4-H264-MEDIUM-30FPS",
			},
		},
	}
	resp, err := client.DeleteMedia(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
	}
}

func TestDescribeAttachedMediaInfo(t *testing.T) {
	client := NewVodClient(nil)
	client.ServiceInfo.Host = HOST_ENDPOINT
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST}
	req := DescribeAttachedMediaInfoRequest{
		MediaKeys: "aa2e88a7e95db9f7,erererer",
	}
	resp, err := client.DescribeAttachedMediaInfo(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestSubmitWorkflow(t *testing.T) {
	client := NewVodClient(nil)
	client.ServiceInfo.Host = HOST_ENDPOINT
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST}
	req := SubmitWorkflowRequest{
		MediaId:    "17c305ac2d3bc999",
		WorkflowId: "3ac7ee5b-0e27-4e35-991f-3facb4fb64e6",
	}
	resp, err := client.SubmitWorkflow(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestProcessMedia(t *testing.T) {
	client := NewVodClient(nil)
	client.ServiceInfo.Host = HOST_ENDPOINT
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST}
	req := ProcessMediaRequest{
		MediaId:      "70805112a18f9e82",
		CallbackArgs: "test",
		TranscodeSets: []TranscodeSet{
			{
				TranscodeTemplateId: "480P_MP4_H264_0",
				URLPath:             "/test.mp4",
				Fps:                 "25",
				Gop:                 "10s",
				WatermarkSets: []WatermarkSet{
					{
						WatermarkTemplateId: "demo_watermark_template_3",
						Height:              "200",
						MarginY:             "0.25",
					},
				},
			},
		},
	}
	resp, err := client.ProcessMedia(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestSubmitMediaProcessJobs(t *testing.T) {
	client := NewVodClient(nil)
	client.ServiceInfo.Host = HOST_ENDPOINT
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST}
	req := SubmitMediaProcessJobsRequest{
		MediaId: "",
		ProcessSet: ProcessSet{
			OperationSets: []OperationSet{
				{
					TemplateId:  "",
					ProcessType: "MediaFeatureAnalysis",
					InputFileSet: InputFileSet{
						Url: "https://s2-10623.kwimgs.com/kos/nlav10623/vision_images/topBannerx1.png",
					},
					OutputFileSet: OutputFileSet{},
					ExtraParams: map[string]string{
						"InputFormat":   "Video",
						"MediaFeatures": "[\"QualityFeature\",\"AestheticsFeature\",\"ContentFeature\",\"AudioFeature\"]",
					},
				},
			},
			CallbackUrl:    "",
			CallbackMethod: "",
			UserData:       "",
		},
	}
	resp, err := client.SubmitMediaProcessJobs(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestDescribeMediaProcessJobs(t *testing.T) {
	client := NewVodClient(nil)
	client.ServiceInfo.Host = HOST_ENDPOINT
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST}
	req := DescribeMediaProcessJobsRequest{
		JobIds: "27eb3ac7-425d-4dfb-91de-769b279b5f06-20220615-0,c7f17271-fa36-48db-94df-e5471f96b09e-20220615-0",
	}
	resp, err := client.DescribeMediaProcessJobs(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestDescribePlayQualityFilters(t *testing.T) {
	client := NewVodClient(nil)
	client.ServiceInfo.Host = HOST_ENDPOINT
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST}
	req := DescribePlayQualityDataSourcesRequest{
		StartTime:   "2023-03-20T16:00:00Z",
		EndTime:     "2023-03-23T15:59:59Z",
		Metric:      "PlayPerformance",
		QueryFilter: []string{"Province", "ISP", "Network"},
	}
	resp, err := client.DescribePlayQualityDataSources(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestDescribePlayQualityOverView(t *testing.T) {
	client := NewVodClient(nil)
	client.ServiceInfo.Host = HOST_ENDPOINT
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST}
	req := DescribePlayQualitySummaryRequest{
		StartTime: "2023-03-20T16:00:00Z",
		EndTime:   "2023-03-23T15:59:59Z",
		Filters: PlayQualityFilterInfo{
			Province:   []string{},
			ISP:        []string{},
			Network:    []string{},
			Platform:   []string{},
			AppVersion: []string{},
			Codec:      []string{},
			Resolution: []string{},
		},
		Metric: "PlayPerformance",
	}
	resp, err := client.DescribePlayQualitySummary(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestDescribePlayQualityDetail(t *testing.T) {
	client := NewVodClient(nil)
	client.ServiceInfo.Host = HOST_ENDPOINT
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST}
	req := DescribePlayQualityDetailRequest{
		StartTime: "2023-03-20T16:00:00Z",
		EndTime:   "2023-03-23T15:59:59Z",
		Filters: PlayQualityFilterInfo{
			Province:   []string{},
			ISP:        []string{},
			Network:    []string{},
			Platform:   []string{},
			AppVersion: []string{},
			Codec:      []string{},
			Resolution: []string{},
		},
		Metric:    "PlayCount",
		Interval:  "5minutes",
		Dimension: []string{"Domain", "Province"},
		Top:       "5",
		Sort:      "AscByAvg",
	}
	resp, err := client.DescribePlayQualityDetail(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestDetectMedia(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   "vod.streamlakeapi.com",
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: "", SecretAccessKey: ""},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := DetectMediaRequest{
		CallbackLink: "https://xxxxxx.xxxxx.com",
		MediaItemSet: MediaItemSet{
			Bucket:       "sl-75abxxxxxxxf19469",
			StoreKey:     "test-xxxxxxxxxxxxx88888.mp4",
			MediaType:    "IMAGE",
			ProcessTypes: []string{"Tag"},
			ClientInfo: ClientInfo{
				TaskId:    "asfbo1bouasndoin1",
				TokenName: "abg1hoasnci",
				Token:     "289ghiusqhoi",
			},
			Text: "风景",
		},
	}
	resp, err := client.DetectMedia(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
	}
}

func TestCreateTranscodeTemplate1(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := CreateTranscodeTemplateRequest{
		TranscodeTemplate: TranscodeTemplate{
			VideoTemplate: VideoTemplate{
				Fps:        20,
				MaxBitrate: 2000,
			},
			AudioTemplate: AudioTemplate{
				Bitrate:    256,
				SampleRate: 44100,
			},
		},
	}
	resp, err := client.CreateTranscodeTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestCreateTranscodeTemplateRemoveAudio(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := CreateTranscodeTemplateRequest{
		TranscodeTemplate: TranscodeTemplate{
			VideoTemplate: VideoTemplate{
				Fps:        20,
				MaxBitrate: 2000,
			},
			RemoveAudio: "true",
		},
	}
	resp, err := client.CreateTranscodeTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestCreateTranscodeTemplate2(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := CreateTranscodeTemplateRequest{
		TranscodeTemplate: TranscodeTemplate{
			Name:        "TranscodeTemplate_12345",
			Description: "test",
			Container:   "hls",
			RemoveAudio: "false",
			VideoTemplate: VideoTemplate{
				Codec:         "libx265",
				Fps:           20,
				MaxBitrate:    3000,
				LongShortMode: "true",
				Width:         50,
				Height:        20,
				Crf:           18,
				Gop:           30,
			},
			AudioTemplate: AudioTemplate{
				Codec:      "mp3",
				Bitrate:    256,
				SampleRate: 44100,
			},
			WatermarkTemplateIds: []string{"db6a7898-e6c4-43d5-9b28-d44a6a1aad1b"},
		},
	}
	resp, err := client.CreateTranscodeTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestUpdateTranscodeTemplate1(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := UpdateTranscodeTemplateRequest{
		TranscodeTemplateId: "e67524fc-fba3-45b7-a040-da6a92ccb787",
		TranscodeTemplate: TranscodeTemplate{
			Name: "transcode_" + "only_test_1",
		},
	}
	resp, err := client.UpdateTranscodeTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestUpdateTranscodeTemplate2(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := UpdateTranscodeTemplateRequest{
		TranscodeTemplateId: "e67524fc-fba3-45b7-a040-da6a92ccb787",
		TranscodeTemplate: TranscodeTemplate{
			TemplateId:  "TemplateId-0",
			Name:        "TranscodeTemplate:" + "test-name-template_124",
			Description: "test",
			Container:   "mp4",
			RemoveAudio: "false",
			VideoTemplate: VideoTemplate{
				Codec:         "libx264",
				Fps:           20,
				MaxBitrate:    2000,
				LongShortMode: "true",
				Width:         0,
				Height:        0,
				Crf:           18,
				Gop:           10,
			},
			AudioTemplate: AudioTemplate{
				Codec:      "mp3",
				Bitrate:    256,
				SampleRate: 44100,
			},
			WatermarkTemplateIds: []string{"db6a7898-e6c4-43d5-9b28-d44a6a1aad1b"},
		},
	}
	resp, err := client.UpdateTranscodeTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestDescribeTranscodeTemplate(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := DescribeTranscodeTemplateRequest{
		TranscodeTemplateId: "272dd39c-bf28-4753-b936-4e7836f80f5c",
	}
	resp, err := client.DescribeTranscodeTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestDescribeTranscodeTemplateIdxxx(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := DescribeTranscodeTemplateRequest{
		TranscodeTemplateId: "xxx",
	}
	_, err := client.DescribeTranscodeTemplate(req)

	assert.Contains(t, err.Error(), "404")
	assert.Contains(t, err.Error(), "ResourceNotFound")
	assert.Contains(t, err.Error(), "TranscodeTemplate does not exist")
}

func TestListTranscodeTemplate1(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := ListTranscodeTemplateRequest{}
	resp, err := client.ListTranscodeTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestListTranscodeTemplate2(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := ListTranscodeTemplateRequest{
		Offset: 4,
		Limit:  45,
	}
	resp, err := client.ListTranscodeTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestListTranscodeTemplateError(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := ListTranscodeTemplateRequest{
		Offset: 4,
		Limit:  51,
	}
	_, err := client.ListTranscodeTemplate(req)

	assert.Contains(t, err.Error(), "400")
	assert.Contains(t, err.Error(), "InvalidArgument")
	assert.Contains(t, err.Error(), "Limit is over 50")
}

func TestDeleteTranscodeTemplate(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := DeleteTranscodeTemplateRequest{
		TranscodeTemplateId: "2c2a7d34-f4aa-48d9-99c0-ecdfec8ec0bc",
	}
	_, err := client.DeleteTranscodeTemplate(req)
	assert.Contains(t, err.Error(), "404")
	assert.Contains(t, err.Error(), "ResourceNotFound")
	assert.Contains(t, err.Error(), "TranscodeTemplate does not exist")
}

func TestCreateWatermarkTemplate1(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := CreateWatermarkTemplateRequest{
		WatermarkTemplate: WatermarkTemplate{
			Name: "watermark_" + "only_test_1",
			Type: "text",
			TextTemplate: TextTemplate{
				Text: "test:" + "only_test_1",
			},
		},
	}
	resp, err := client.CreateWatermarkTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestCreateWatermarkTemplate2(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := CreateWatermarkTemplateRequest{
		WatermarkTemplate: WatermarkTemplate{
			Name: "watermark-" + "only_test_1",
			Type: "image",
			ImageTemplate: ImageTemplate{
				Resource: Resource{
					Bucket: "mediacloud-streamlake-app_video",
					Object: "cdn_test_pic.png",
				},
			},
		},
	}
	resp, err := client.CreateWatermarkTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestCreateWatermarkTemplate3(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := CreateWatermarkTemplateRequest{
		WatermarkTemplate: WatermarkTemplate{
			TemplateId:    "WatermarkTemplate-0",
			Name:          "watermark-" + "only_test_name",
			Description:   "test",
			Type:          "text",
			ReferPosition: "topLeft",
			MarginX:       "0.43",
			MarginY:       "0.02",
			TextTemplate: TextTemplate{
				FontType:  "SourceHanSans",
				Text:      "test:" + "test-name-template_124",
				FontSize:  40,
				FontColor: "#FF0000",
			},
		},
	}
	resp, err := client.CreateWatermarkTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestCreateWatermarkTemplate4(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := CreateWatermarkTemplateRequest{
		WatermarkTemplate: WatermarkTemplate{
			TemplateId:    "WatermarkTemplate-0",
			Name:          "watermark:" + "test-name-template_124",
			Description:   "test",
			Type:          "image",
			ReferPosition: "topLeft",
			MarginX:       "0.99999",
			MarginY:       "0.99999999",
			ImageTemplate: ImageTemplate{
				Resource: Resource{
					Bucket: "mediacloud-streamlake-app_video",
					Object: "cdn_test_pic.png",
				},
				Width:  "0.9999",
				Height: "0.9999",
			},
		},
	}
	resp, err := client.CreateWatermarkTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestUpdateWatermarkTemplate1(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := UpdateWatermarkTemplateRequest{
		WatermarkTemplateId: "db6a7898-e6c4-43d5-9b28-d44a6a1aad1b",
		WatermarkTemplate: WatermarkTemplate{
			Name: "watermark---" + "only_test_1",
			TextTemplate: TextTemplate{
				Text: "test___" + "only_test_1",
			},
		},
	}
	resp, err := client.UpdateWatermarkTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestUpdateWatermarkTemplate2(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := UpdateWatermarkTemplateRequest{
		WatermarkTemplateId: "db6a7898-e6c4-43d5-9b28-d44a6a1aad1b",
		WatermarkTemplate: WatermarkTemplate{
			TemplateId:    "WatermarkTemplate-0",
			Name:          "watermark_" + "test-name-template_124",
			Description:   "test",
			Type:          "text",
			ReferPosition: "topLeft",
			MarginX:       "0.5",
			MarginY:       "0.3",
			TextTemplate: TextTemplate{
				FontType:  "SourceHanSans",
				Text:      "test:" + "test-name-template_124",
				FontSize:  40,
				FontColor: "#FF0000",
			},
		},
	}
	_, err := client.UpdateWatermarkTemplate(req)

	assert.Contains(t, err.Error(), "400")
	assert.Contains(t, err.Error(), "InvalidArgument")
	assert.Contains(t, err.Error(), "Can't change type in watermarkTemplate")
}

func TestUpdateWatermarkTemplate3(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := UpdateWatermarkTemplateRequest{
		WatermarkTemplateId: "db6a7898-e6c4-43d5-9b28-d44a6a1aad1b",
		WatermarkTemplate: WatermarkTemplate{
			TemplateId:    "WatermarkTemplate-0",
			Name:          "watermark:" + "test-name-template_124",
			Description:   "test",
			Type:          "image",
			ReferPosition: "topLeft",
			MarginX:       "0",
			MarginY:       "0",
			ImageTemplate: ImageTemplate{
				Resource: Resource{
					Bucket: "mediacloud-streamlake-app_video",
					Object: "cdn_test_pic.png",
				},
				Width:  "0.1",
				Height: "0.1",
			},
		},
	}
	resp, err := client.UpdateWatermarkTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestDescribeWatermarkTemplate(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := DescribeWatermarkTemplateRequest{
		WatermarkTemplateId: "db6a7898-e6c4-43d5-9b28-d44a6a1aad1b",
	}
	resp, err := client.DescribeWatermarkTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestListWatermarkTemplate1(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := ListWatermarkTemplateRequest{}
	resp, err := client.ListWatermarkTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestListWatermarkTemplate2(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := ListWatermarkTemplateRequest{
		Offset: 0,
		Limit:  50,
	}
	resp, err := client.ListWatermarkTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestDeleteWatermarkTemplate(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := DeleteWatermarkTemplateRequest{
		WatermarkTemplateId: "8eea6738-6505-4f5a-b625-8ce6d07c5ed7",
	}
	resp, err := client.DeleteWatermarkTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
	}
}

func TestCreateSnapshotTemplate1(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := CreateSnapshotTemplateRequest{
		TemplateType: "SampleSnapshot",
		SampleSnapshotTemplate: SampleSnapshotTemplate{
			Interval: 2,
			Height:   50,
		},
	}
	resp, err := client.CreateSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestCreateSnapshotTemplateNotSetWidthAndHeight(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := CreateSnapshotTemplateRequest{
		TemplateType: "SampleSnapshot",
		SampleSnapshotTemplate: SampleSnapshotTemplate{
			Interval: 1,
		},
	}
	_, err := client.CreateSnapshotTemplate(req)
	assert.Contains(t, err.Error(), "400")
	assert.Contains(t, err.Error(), "InvalidArgument")
	assert.Contains(t, err.Error(), "Width and Height can not both be null or zero")
}

func TestCreateSnapshotTemplate2(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := CreateSnapshotTemplateRequest{
		TemplateType: "SnapshotByTimeOffset",
		SnapshotByTimeOffsetTemplate: SnapshotByTimeOffsetTemplate{
			//Name: "SnapshotByTimeOffsetTemplate:" + "only_test_1",
			Name: strings.ReplaceAll("name-SnapshotByTimeOffset", " ", ""),
		},
	}
	resp, err := client.CreateSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestCreateSnapshotTemplateSnapshotByTimeOffsetTemplateOptional(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := CreateSnapshotTemplateRequest{
		TemplateType:                 "SnapshotByTimeOffset",
		SnapshotByTimeOffsetTemplate: SnapshotByTimeOffsetTemplate{},
	}
	resp, err := client.CreateSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestCreateSnapshotTemplate3(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := CreateSnapshotTemplateRequest{
		TemplateType: "ImageSprite",
		ImageSpriteTemplate: ImageSpriteTemplate{
			Name:           "ImageSpriteTemplate-" + "only_test_1",
			SampleInterval: 1,
			RowCount:       3,
			ColumnCount:    5,
			Width:          200,
		},
	}
	resp, err := client.CreateSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestCreateSnapshotTemplateNotSetWidthHeight(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := CreateSnapshotTemplateRequest{
		TemplateType: "ImageSprite",
		ImageSpriteTemplate: ImageSpriteTemplate{
			Name:           "ImageSpriteTemplate-" + "only_test_1",
			SampleInterval: 1,
			RowCount:       3,
			ColumnCount:    5,
		},
	}
	_, err := client.CreateSnapshotTemplate(req)
	assert.Contains(t, err.Error(), "400")
	assert.Contains(t, err.Error(), "InvalidArgument")
	assert.Contains(t, err.Error(), "Width and Height can not both be null or zero")
}

func TestCreateSnapshotTemplate4(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := CreateSnapshotTemplateRequest{
		TemplateType: "SampleSnapshot",
		SampleSnapshotTemplate: SampleSnapshotTemplate{
			Name:        "only_test_SampleSnapshot",
			Description: "test-name-template_124",
			SampleType:  "Time",
			Interval:    2,
			Format:      "png",
			Width:       200,
			Height:      400,
			OffsetTime:  5,
		},
	}
	resp, err := client.CreateSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestCreateSnapshotTemplate5(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := CreateSnapshotTemplateRequest{
		TemplateType: "SnapshotByTimeOffset",
		SnapshotByTimeOffsetTemplate: SnapshotByTimeOffsetTemplate{
			Name:        strings.ReplaceAll("name-SnapshotByTimeOffset", " ", ""),
			Description: "test-name-template_124",
			Format:      "png",
			Width:       0,
			Height:      0,
			OffsetTime:  0,
		},
	}
	resp, err := client.CreateSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestCreateSnapshotTemplate6(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := CreateSnapshotTemplateRequest{
		TemplateType: "ImageSprite",
		ImageSpriteTemplate: ImageSpriteTemplate{
			SnapshotTemplateId: "SnapshotTemplateId-0",
			Name:               strings.ReplaceAll("name-ImageSpriteTemplate", " ", ""),
			Description:        "test-name-template_124",
			SampleType:         "Time",
			SampleInterval:     1,
			RowCount:           3,
			ColumnCount:        5,
			Width:              200,
			Height:             420,
			Format:             "png",
		},
	}
	resp, err := client.CreateSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestUpdateSnapshotTemplate1(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := UpdateSnapshotTemplateRequest{
		TemplateType:       "SampleSnapshot",
		SnapshotTemplateId: "80d5cce2-8877-4662-9e50-f268baeaa21d",
		SampleSnapshotTemplate: SampleSnapshotTemplate{
			Name: "SampleSnapshot:" + "only_test_1",
		},
	}
	resp, err := client.UpdateSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestUpdateSnapshotTemplate2(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := UpdateSnapshotTemplateRequest{
		TemplateType:       "SnapshotByTimeOffset",
		SnapshotTemplateId: "9387156c-2d0c-4818-8cbc-f61a5d62ffb9",
		SnapshotByTimeOffsetTemplate: SnapshotByTimeOffsetTemplate{
			Name: "SnapshotByTimeOffsetTemplate:" + "only_test_1",
		},
	}
	_, err := client.UpdateSnapshotTemplate(req)
	assert.Contains(t, err.Error(), "404")
	assert.Contains(t, err.Error(), "ResourceNotFound")
	assert.Contains(t, err.Error(), "Template is not found")
}

func TestUpdateSnapshotTemplate3(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := UpdateSnapshotTemplateRequest{
		TemplateType:       "ImageSprite",
		SnapshotTemplateId: "0a00576f-05ce-48d9-b453-3a5f1f72e235",
		ImageSpriteTemplate: ImageSpriteTemplate{
			Name: "ImageSprite:" + "only_test_1",
		},
	}
	resp, err := client.UpdateSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestUpdateSnapshotTemplate4(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := UpdateSnapshotTemplateRequest{
		TemplateType:       "SampleSnapshot",
		SnapshotTemplateId: "f0ecfb68-5c94-4655-96b3-a045f67c7deb",
		SampleSnapshotTemplate: SampleSnapshotTemplate{
			SnapshotTemplateId: "SnapshotTemplateId-0",
			Name:               strings.ReplaceAll("name-SampleSnapshot", " ", ""),
			Description:        "test-name-template_124",
			SampleType:         "Time",
			Interval:           1,
			Format:             "jpeg",
			Width:              0,
			Height:             0,
			OffsetTime:         0,
		},
	}
	resp, err := client.UpdateSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestUpdateSnapshotTemplate5(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := UpdateSnapshotTemplateRequest{
		TemplateType:       "SnapshotByTimeOffset",
		SnapshotTemplateId: "9387156c-2d0c-4818-8cbc-f61a5d62ffb9",
		SnapshotByTimeOffsetTemplate: SnapshotByTimeOffsetTemplate{
			SnapshotTemplateId: "SnapshotTemplateId-0",
			Description:        "test-name-template_124",
			Format:             "png",
			Width:              0,
			Height:             0,
			OffsetTime:         0,
		},
	}
	resp, err := client.UpdateSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestUpdateSnapshotTemplate6(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := UpdateSnapshotTemplateRequest{
		TemplateType:       "ImageSprite",
		SnapshotTemplateId: "0a00576f-05ce-48d9-b453-3a5f1f72e235",
		ImageSpriteTemplate: ImageSpriteTemplate{
			SnapshotTemplateId: "SnapshotTemplateId-0",
			Name:               strings.ReplaceAll("name-ImageSpriteTemplate", " ", ""),
			Description:        "test-name-template_124",
			SampleType:         "Time",
			SampleInterval:     1,
			RowCount:           3,
			ColumnCount:        5,
			Width:              0,
			Height:             0,
			Format:             "png",
		},
	}
	resp, err := client.UpdateSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestDescribeSnapshotTemplate1(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := DescribeSnapshotTemplateRequest{
		TemplateType:       "SampleSnapshot",
		SnapshotTemplateId: "80d5cce2-8877-4662-9e50-f268baeaa21d",
	}
	resp, err := client.DescribeSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestDescribeSnapshotTemplate2(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := DescribeSnapshotTemplateRequest{
		TemplateType:       "SnapshotByTimeOffset",
		SnapshotTemplateId: "a5d62319-3f8c-4431-8cca-24e562a2b023",
	}
	resp, err := client.DescribeSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestDescribeSnapshotTemplateAnotherAK(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := DescribeSnapshotTemplateRequest{
		TemplateType:       "SnapshotByTimeOffset",
		SnapshotTemplateId: "b599e91c-1b76-4837-a06e-d400d3f9f428",
	}
	_, err := client.DescribeSnapshotTemplate(req)
	assert.Contains(t, err.Error(), "404")
	assert.Contains(t, err.Error(), "ResourceNotFound")
	assert.Contains(t, err.Error(), "Template is not found")
}

func TestDescribeSnapshotTemplate3(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := DescribeSnapshotTemplateRequest{
		TemplateType:       "ImageSprite",
		SnapshotTemplateId: "0a00576f-05ce-48d9-b453-3a5f1f72e235",
	}
	resp, err := client.DescribeSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestListSnapshotTemplate1(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := ListSnapshotTemplateRequest{
		TemplateType: "SampleSnapshot",
	}
	resp, err := client.ListSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestListSnapshotTemplate2(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := ListSnapshotTemplateRequest{
		TemplateType: "SnapshotByTimeOffset",
	}
	resp, err := client.ListSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestListSnapshotTemplate3(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := ListSnapshotTemplateRequest{
		TemplateType: "ImageSprite",
	}
	resp, err := client.ListSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestListSnapshotTemplate4(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := ListSnapshotTemplateRequest{
		TemplateType: "ImageSprite",
		Offset:       0,
		Limit:        30,
	}
	resp, err := client.ListSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestDeleteSnapshotTemplate1(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := DeleteSnapshotTemplateRequest{
		TemplateType:       "SampleSnapshot",
		SnapshotTemplateId: "f0ecfb68-5c94-4655-96b3-a045f67c7deb",
	}
	resp, err := client.DeleteSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestDeleteSnapshotTemplate2(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := DeleteSnapshotTemplateRequest{
		TemplateType:       "SnapshotByTimeOffset",
		SnapshotTemplateId: "9387156c-2d0c-4818-8cbc-f61a5d62ffb9",
	}
	_, err := client.DeleteSnapshotTemplate(req)
	assert.Contains(t, err.Error(), "404")
	assert.Contains(t, err.Error(), "ResourceNotFound")
	assert.Contains(t, err.Error(), "Template is not found")
}

func TestDeleteSnapshotTemplate3(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := DeleteSnapshotTemplateRequest{
		TemplateType:       "ImageSprite",
		SnapshotTemplateId: "0a00576f-05ce-48d9-b453-3a5f1f72e235",
	}
	resp, err := client.DeleteSnapshotTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestCreateWorkflowTemplate(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := CreateWorkflowTemplateRequest{
		MediaProcessWorkflowTemplate: MediaProcessWorkflowTemplate{
			WorkflowName: "WorkflowName_" + "only_test_1",
			Description:  "test-name-template_124",
			TranscodeTasks: []TranscodeTask{{
				TemplateId: "e67524fc-fba3-45b7-a040-da6a92ccb787",
			}},
			SampleSnapshotTasks: []SampleSnapshotTask{{
				TemplateId: "8a511fdd-eb6b-4c34-8d2d-dc5ce4cbdc06",
			}},
			SnapshotByTimeOffsetTasks: []SnapshotByTimeOffsetTask{{
				TemplateId: "5d6029d6-ed55-4eec-8704-b7153a635fae",
			}},
			ImageSpriteTasks: []ImageSpriteTask{{
				TemplateId: "b2f8ccfc-2fb1-4b7f-90f0-9c21255a5345",
			}},
		},
	}
	resp, err := client.CreateWorkflowTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestCreateWorkflowTemplateOnlySnapshot(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := CreateWorkflowTemplateRequest{
		MediaProcessWorkflowTemplate: MediaProcessWorkflowTemplate{
			WorkflowName: "WorkflowName_" + "only_test_3",
			Description:  "test-name-template_124",
			SampleSnapshotTasks: []SampleSnapshotTask{{
				TemplateId: "8a511fdd-eb6b-4c34-8d2d-dc5ce4cbdc06",
			}},
		},
	}
	resp, err := client.CreateWorkflowTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestCreateWorkflowTemplateOnlyError(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := CreateWorkflowTemplateRequest{
		MediaProcessWorkflowTemplate: MediaProcessWorkflowTemplate{
			WorkflowName: "WorkflowName_" + "only_test_3",
			Description:  "test-name-template_124",
		},
	}
	_, err := client.CreateWorkflowTemplate(req)
	assert.Contains(t, err.Error(), "400")
	assert.Contains(t, err.Error(), "InvalidArgument")
	assert.Contains(t, err.Error(), "MediaProcessTemplates are all empty")
}

func TestUpdateWorkflowTemplate(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := UpdateWorkflowTemplateRequest{
		WorkflowId: "3feeae66-de25-4b8c-8686-0064092a331d",
		MediaProcessWorkflowTemplate: MediaProcessWorkflowTemplate{
			WorkflowName: "WorkflowName_" + "only_test_1",
			Description:  "test-name-template_124",
			TranscodeTasks: []TranscodeTask{{
				TemplateId: "e67524fc-fba3-45b7-a040-da6a92ccb787",
			}},
			SampleSnapshotTasks: []SampleSnapshotTask{{
				TemplateId: "8a511fdd-eb6b-4c34-8d2d-dc5ce4cbdc06",
			}},
			SnapshotByTimeOffsetTasks: []SnapshotByTimeOffsetTask{{
				TemplateId: "5d6029d6-ed55-4eec-8704-b7153a635fae",
			}},
			ImageSpriteTasks: []ImageSpriteTask{{
				TemplateId: "b2f8ccfc-2fb1-4b7f-90f0-9c21255a5345",
			}},
		},
	}
	_, err := client.UpdateWorkflowTemplate(req)
	assert.Contains(t, err.Error(), "404")
	assert.Contains(t, err.Error(), "ResourceNotFound")
	assert.Contains(t, err.Error(), "WorkflowTemplate does not exist")
}

func TestUpdateWorkflowTemplateNormal(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := UpdateWorkflowTemplateRequest{
		WorkflowId: "8a511fdd-eb6b-4c34-8d2d-dc5ce4cbdc06",
		MediaProcessWorkflowTemplate: MediaProcessWorkflowTemplate{
			WorkflowName: "WorkflowName_" + "only_test_1",
			Description:  "test-name-template_124",
			TranscodeTasks: []TranscodeTask{{
				TemplateId: "e67524fc-fba3-45b7-a040-da6a92ccb787",
			}},
			SampleSnapshotTasks: []SampleSnapshotTask{{
				TemplateId: "8a511fdd-eb6b-4c34-8d2d-dc5ce4cbdc06",
			}},
			SnapshotByTimeOffsetTasks: []SnapshotByTimeOffsetTask{{
				TemplateId: "5d6029d6-ed55-4eec-8704-b7153a635fae",
			}},
			ImageSpriteTasks: []ImageSpriteTask{{
				TemplateId: "b2f8ccfc-2fb1-4b7f-90f0-9c21255a5345",
			}},
		},
	}
	_, err := client.UpdateWorkflowTemplate(req)
	assert.Contains(t, err.Error(), "404")
	assert.Contains(t, err.Error(), "ResourceNotFound")
	assert.Contains(t, err.Error(), "WorkflowTemplate does not exist")
}

func TestListWorkflowTemplate1(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := ListWorkflowTemplateRequest{}
	resp, err := client.ListWorkflowTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestListWorkflowTemplate2(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := ListWorkflowTemplateRequest{
		Names:  []string{""},
		Offset: 1,
		Limit:  30,
	}
	resp, err := client.ListWorkflowTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestDeleteWorkflowTemplate(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := DeleteWorkflowTemplateRequest{
		WorkflowId: "3feeae66-de25-4b8c-8686-0064092a331d",
	}
	resp, err := client.DeleteWorkflowTemplate(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		t.Logf("got response data: %+v", resp.ResponseData)
	}
}

func TestSubmitWorkflow2(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := SubmitWorkflowRequest{
		MediaId:    MEDIA_ID,
		WorkflowId: "ad6e4768-29fd-414c-a56b-37fe36e59ac1",
	}
	resp, err := client.SubmitWorkflow(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestDescribeTaskDetail(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)

	req := DescribeTaskDetailRequest{
		TaskId: "d6f1f774128d04c0f6a2ef72f32932b4",
	}
	resp, err := client.DescribeTaskDetail(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
		json, _ := json.Marshal(resp.ResponseData)
		t.Logf("got response data: %+s", json)
	}
}

// 获取上传凭证样例
func TestApplyUploadInfo(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type":  []string{"application/json"},
			"trace-context": []string{"{\"laneId\":\"PRT.StreamLake\"}"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)
	req := ApplyUploadInfoRequest{
		FilePath: "test.mp4",
		Format:   "mp4",
	}
	resp, err := client.ApplyUploadInfo(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

// 确认上传样例
func TestCommitUploadInfo(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type":  []string{"application/json"},
			"trace-context": []string{"{\"laneId\":\"PRT.StreamLake\"}"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)
	req := CommitUploadRequest{
		SessionKey: "xxxxxxxxxx",
	}
	resp, err := client.CommitUpload(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestGeneratePreSignedUrl(t *testing.T) {
	Endpoint := "https://kms-cn-beijing.streamlakeapi.com"
	AccessKey := ACCESS_KEY_TEST
	SecretKey := SECRET_KEY_TEST
	bucketName := "zhf-test-0509-02"
	objectKey := "221.png"
	expiration := 3600

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("cn-beijing"),
		Credentials: credentials.NewStaticCredentials(AccessKey, SecretKey, ""),
		Endpoint:    aws.String(Endpoint),
	})
	if err != nil {
		log.Printf("GetObject error: %v", err)
		return
	}
	//使用
	svc := s3.New(sess)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})

	presignedURL, _ := req.Presign(time.Duration(expiration) * time.Second)
	fmt.Println(presignedURL)
}
func TestUpdateDanamkuAuditResult(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)
	request := UpdateDanmakuAuditResultRequest{
		Status:    "Pass",
		DanmakuId: 301903,
		//DanmakuIds: []int64{301903},
	}
	resp, err := client.updateDanmakuPreAuditResult(request)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestListDanmakuPreAuditNeccesary(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Region: "cn-beijing",
		Scheme: "https",
		Host:   HOST_ENDPOINT,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		ProductName: "vod",
		Credentials: base.Credentials{AccessKey: ACCESS_KEY_TEST, SecretAccessKey: SECRET_KEY_TEST},
	}
	client := NewVodClientV2(nil, serviceInfo)
	request := ListDanmakuPreAuditRequest{}

	resp, err := client.listDanmakuPreAudit(request)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}

}
