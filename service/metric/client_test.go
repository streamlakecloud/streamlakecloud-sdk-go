package metric

import (
	"bufio"
	"compress/gzip"
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
func TestDescribeNCdnLogs(t *testing.T) {
	client := NewMetricClient(nil)
	client.ServiceInfo.Host = DEMO_HOST_STAGING
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: DEMO_TEST_ACCESS_KEY}
	req := DescribeNCdnLogsRequest{
		StartTime: "2024-11-21T16:00:00Z",
		EndTime:   "2024-11-22T23:59:59Z",
	}
	resp, err := client.DescribeNCdnLogs(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response, logInfo: %s", resp.ResponseData.LogInfos)
	}
}

func TestDownloadNCdnLog(t *testing.T) {
	client := NewMetricClient(nil)
	client.ServiceInfo.Host = DEMO_HOST_STAGING
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: DEMO_TEST_ACCESS_KEY}
	req := DownloadNCdnLogRequest{
		LogFileName: "202411221825.gz",
	}
	resp, err := client.DownloadNCdnLog(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("Downloaded Headers %v", resp.Header)
		defer resp.Body.Close()

		// read the response body, gunzip it
		fz, err := gzip.NewReader(resp.Body)
		if err != nil {
			t.Fatalf("%e", err)
		}
		defer fz.Close()

		// read first 10 lines
		scanner := bufio.NewScanner(fz)
		for i := 0; i < 10 && scanner.Scan(); i++ {
			t.Logf("Line %d: %s", i+1, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			t.Fatalf("%e", err)
		}
	}
}
