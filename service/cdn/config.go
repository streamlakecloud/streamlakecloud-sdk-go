package cdn

import (
	"net/http"
	"net/url"

	"github.com/streamlakecloud/streamlakecloud-sdk-go/base"
)

var ServiceInfo = base.ServiceInfo{
	Region: "cn-beijing",
	Scheme: "https",
	Host:   "vod.streamlakeapi.com",
	Header: http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
	},
	Credentials: base.Credentials{},
}

var ApiList = map[string]*base.ApiInfo{
	// CDN
	"DescribeDomainRealTimeOriginData": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DescribeDomainRealTimeOriginData"},
		},
	},
	"DescribeDomainRealTimeCdnData": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DescribeDomainRealTimeCdnData"},
		},
	},
	"DescribeCdnLogs": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DescribeCdnLogs"},
		},
	},
	"PreloadObjectCaches": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"PreloadObjectCaches"},
		},
	},
	"RefreshObjectCaches": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"RefreshObjectCaches"},
		},
	},
	"DescribeRefreshTasks": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DescribeRefreshTasks"},
		},
	},
	"PushPCDNObjectCache": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"PushPCDNObjectCache"},
		},
	},
}
