package metric

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/streamlakecloud/streamlakecloud-sdk-go/base"
)

var (
	DEMO_HOST_STAGING    = os.Getenv("STREAMLAKE_VOD_GO_SDK_ENDPOINT")
	DEMO_TEST_ACCESS_KEY = os.Getenv("STREAMLAKE_VOD_GO_SDK_ACCESS_KEY")
)

func TestDescribeStorageUsageData(t *testing.T) {
	client := NewMetricClient(nil)
	client.ServiceInfo.Host = DEMO_HOST_STAGING
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: DEMO_TEST_ACCESS_KEY}
	req := DescribeStorageUsageDataRequest{
		StartTime: "2022-05-01T14:00:00Z",
		EndTime:   "2022-05-10T23:00:00Z",
		Region:    "cn-beijing",
		Bucket:    "def",
		Interval:  "day",
	}
	resp, err := client.DescribeStorageUsageData(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestRawOpenAPIVodClient(t *testing.T) {
	serviceInfo := base.ServiceInfo{
		Host:        DEMO_HOST_STAGING,
		Credentials: base.Credentials{AccessKey: DEMO_TEST_ACCESS_KEY},
	}
	apiInfoMap := ApiList
	c := base.NewClient(nil, serviceInfo, apiInfoMap)
	req := DescribeStorageUsageDataRequest{
		StartTime: "2022-05-01T14:00:00Z",
		EndTime:   "2022-05-10T23:00:00Z",
		Region:    "cn-beijing",
		Bucket:    "def",
		Interval:  "day",
	}
	resp, e := c.PostForAPIWithRequest("DescribeStorageUsageData", req)
	if e != nil {
		t.Fatalf("%e", e)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("%e", e)
		}
		t.Logf("%s", body)
	}
}
