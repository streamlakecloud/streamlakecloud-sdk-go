package vod

import (
	"os"
	"testing"

	"github.com/streamlakecloud/streamlakecloud-sdk-go/base"
)

var (
	DEMO_HOST_STAGING    = os.Getenv("STREAMLAKE_VOD_GO_SDK_ENDPOINT")
	DEMO_TEST_ACCESS_KEY = os.Getenv("STREAMLAKE_VOD_GO_SDK_ACCESS_KEY2")
	DEMO_TEST_SECRET_KEY = os.Getenv("STREAMLAKE_VOD_GO_SDK_SECRET_KEY")
)

func TestApplyUploadToken(t *testing.T) {
	client := NewVodClient(nil)
	client.ServiceInfo.Host = DEMO_HOST_STAGING
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: DEMO_TEST_ACCESS_KEY, SecretAccessKey: DEMO_TEST_SECRET_KEY}
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
	client.ServiceInfo.Host = DEMO_HOST_STAGING
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: DEMO_TEST_ACCESS_KEY, SecretAccessKey: DEMO_TEST_SECRET_KEY}
	req := FetchUploadRequest{
		URLSets: []FetchUploadURLSet{
			{
				SourceURL:    "https://cdnfile.corp.kuaishou.com/kc/files/a/post-test-materials/big_buck_bunny_720p_5mb.mp4",
				CallbackArgs: "test",
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
	client.ServiceInfo.Host = DEMO_HOST_STAGING
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: DEMO_TEST_ACCESS_KEY, SecretAccessKey: DEMO_TEST_SECRET_KEY}
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
	client.ServiceInfo.Host = DEMO_HOST_STAGING
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: DEMO_TEST_ACCESS_KEY, SecretAccessKey: DEMO_TEST_SECRET_KEY}
	req := FetchStoreRequest{
		URLSets: []FetchStoreURLSet{
			{
				SourceURL:    "https://cdnfile.corp.kuaishou.com/kc/files/a/post-test-materials/big_buck_bunny_720p_5mb.mp4",
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

func TestDescribeMediaInfo(t *testing.T) {
	client := NewVodClient(nil)
	client.ServiceInfo.Host = DEMO_HOST_STAGING
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: DEMO_TEST_ACCESS_KEY, SecretAccessKey: DEMO_TEST_SECRET_KEY}
	req := DescribeMediaInfoRequest{
		MediaId: "17c305ac2d3bc999", // https://video-store.staging.kuaishou.com//readVideoAsset/1ac4beaac9750903
	}
	resp, err := client.DescribeMediaInfo(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestDescribeAttachedMediaInfo(t *testing.T) {
	client := NewVodClient(nil)
	client.ServiceInfo.Host = DEMO_HOST_STAGING
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: DEMO_TEST_ACCESS_KEY, SecretAccessKey: DEMO_TEST_SECRET_KEY}
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
	client.ServiceInfo.Host = DEMO_HOST_STAGING
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: DEMO_TEST_ACCESS_KEY, SecretAccessKey: DEMO_TEST_SECRET_KEY}
	req := SubmitWorkflowRequest{
		MediaId:    "17c305ac2d3bc999", // https://video-store.staging.kuaishou.com//readVideoAsset/1ac4beaac9750903
		WorkflowId: "workflow_h265_mp4_0",
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
	client.ServiceInfo.Host = DEMO_HOST_STAGING
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: DEMO_TEST_ACCESS_KEY, SecretAccessKey: DEMO_TEST_SECRET_KEY}
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

func TestDescribeMediaProcessJobs(t *testing.T) {
	client := NewVodClient(nil)
	client.ServiceInfo.Host = DEMO_HOST_STAGING
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: DEMO_TEST_ACCESS_KEY, SecretAccessKey: DEMO_TEST_SECRET_KEY}
	req := DescribeMediaProcessJobsRequest{
		JobIds: "27eb3ac7-425d-4dfb-91de-769b279b5f06-20220615-0,c7f17271-fa36-48db-94df-e5471f96b09e-20220615-0", // https://video-store.staging.kuaishou.com//readVideoAsset/1ac4beaac9750903
	}
	resp, err := client.DescribeMediaProcessJobs(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}
