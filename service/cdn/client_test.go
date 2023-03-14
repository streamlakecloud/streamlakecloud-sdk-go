package cdn

import (
	"os"
	"testing"

	"github.com/streamlakecloud/streamlakecloud-sdk-go/base"
)

var (
	DEMO_HOST_STAGING      = os.Getenv("STREAMLAKE_VOD_GO_SDK_ENDPOINT")
	DEMO_HOST_ONLINE       = os.Getenv("STREAMLAKE_VOD_GO_SDK_ENDPOINT_ONLINE")
	DEMO_TEST_ACCESS_KEY   = os.Getenv("STREAMLAKE_VOD_GO_SDK_ACCESS_KEY")
	DEMO_ONLINE_ACCESS_KEY = os.Getenv("STREAMLAKE_VOD_GO_SDK_ACCESS_KEY_ONLINE")
)

func TestDescribeCdnLogs(t *testing.T) {
	client := NewCDNClient(nil)
	client.ServiceInfo.Host = DEMO_HOST_STAGING
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: DEMO_TEST_ACCESS_KEY}
	req := DescribeCdnLogsRequest{
		StartTime:  "2022-06-27 00:00:00",
		EndTime:    "2022-06-27 01:00:00",
		DomainName: "vdn5.vzuu.com",
	}
	resp, err := client.DescribeCdnLogs(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestPreloadObjectCaches(t *testing.T) {
	client := NewCDNClient(nil)
	client.ServiceInfo.Host = DEMO_HOST_ONLINE
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: DEMO_ONLINE_ACCESS_KEY}
	req := PreloadObjectCachesRequest{
		ObjectPath: "http://slplay.streamlake.cn/FvPO-JwwhdMEV_iv5Y_iEZGFV03SrLS1bIuOACLK4KU",
	}
	resp, err := client.PreloadObjectCaches(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestRefreshObjectCaches(t *testing.T) {
	client := NewCDNClient(nil)
	client.ServiceInfo.Host = DEMO_HOST_ONLINE
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: DEMO_ONLINE_ACCESS_KEY}
	req := RefreshObjectCachesRequest{
		ObjectPath: "http://slplay.streamlake.cn/FvPO-JwwhdMEV_iv5Y_iEZGFV03SrLS1bIuOACLK4KU",
	}
	resp, err := client.RefreshObjectCaches(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestDescribeRefreshTasks(t *testing.T) {
	client := NewCDNClient(nil)
	client.ServiceInfo.Host = DEMO_HOST_ONLINE
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: DEMO_ONLINE_ACCESS_KEY}
	req := DescribeRefreshTasksRequest{
		TaskId:   "20220714_a00783c3-79a9-4e19-80db-b365d881027e",
		Limit:    1,
		Offset:   0,
		TaskType: "preload",
	}
	resp, err := client.DescribeRefreshTasks(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestGetCdnData(t *testing.T) {
	client := NewCDNClient(nil)
	client.ServiceInfo.Host = "streamlake-api.staging.kuaishou.com"
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: "8dcf3df43316434080a179a7a573054a", SecretAccessKey: "21025fbc4cdd4a9cb7eace76c3449017"}

	req := GetDomainRealTimeCdnRequest{
		StartTime: "2022-08-29T01:58:00Z",
		EndTime:   "2022-08-29T02:02:00Z",
		Interval:  "1minute",
	}

	resp, err := client.GetDomainRealTimeCdnData(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestGetOriginData(t *testing.T) {
	client := NewCDNClient(nil)
	client.ServiceInfo.Host = "streamlake-api.staging.kuaishou.com"
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: "8dcf3df43316434080a179a7a573054a", SecretAccessKey: "21025fbc4cdd4a9cb7eace76c3449017"}

	req := GetDomainRealTimeOriginRequest{
		StartTime: "2022-08-29T01:58:00Z",
		EndTime:   "2022-08-29T02:02:00Z",
		Interval:  "1minute",
	}

	resp, err := client.GetDomainRealTimeOriginData(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v, data: %+v", resp.ResponseMeta, resp.ResponseData)
	}
}

func TestPushPCDNCache(t *testing.T) {
	client := NewCDNClient(nil)
	client.ServiceInfo.Host = "vod.streamlakeapi.com"
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: "f31f5838ba104040b2f23a51a0c52e84", SecretAccessKey: "69cb6facdf36486d9f7b7cdf943f1fb5"}
	client.ServiceInfo.ProductName = "vod"
	client.ServiceInfo.Header.Set("Content-Type", "application/json")

	req := PushPCDNObjectCacheRequest{
		FileSets: []PCDNFileSet{
			{
				URLPath:  "http://d1-pcdn.a.kwimgs.com/upic/2023/02/15/11/BMjAyMzAyMTUxMTQ0MDRfMTQ3NTgwOTg3N185NjI1OTYzNjM3M18xXzM=_vnushighv2_B3aa492493a1ce82b7822b4467c1734ed.mp4",
				FileSize: 740459,
				Crc32:    "7d4a8656",
				BlockCrc32: []string{"4e60a6c2", "dd10a0df", "d76fc281", "0779122a", "b8766665",
					"baa94ba2", "a807440f", "ca22eef3", "0fd7050a", "fa276a52", "60ed407d", "36d62de0", "037e9e06", "eafb115f",
					"51cc1d5e", "28918d04", "629be459", "6d080dfe", "bc797a28", "9c165c4a", "08951ceb", "dd839c7b", "9e1f16f8",
					"19fb74e1", "d6aab174", "045bdae2", "683e8853", "ebc5bf89", "b2865644", "c4ac63cd", "f4dda727", "eb057dad",
					"d1d79ba5", "d052828c", "3ee1b501", "79e3646d", "b0379ecb", "ce3dc85e", "2b36a44d", "18669fa5", "6455320b",
					"c13dfa11", "33b29b89", "b6be0f95", "c00cfa45", "3b514fa1"},
				Qps: 10,
			},
		},
	}
	resp, err := client.PushPCDNCache(req)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("got response meta: %+v", resp.ResponseMeta)
	}
}
