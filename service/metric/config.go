package metric

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
	// Metrics
	"DescribeDomainUsageData": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DescribeDomainUsageData"},
		},
	},

	"DescribeStorageUsageData": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DescribeStorageUsageData"},
		},
	},

	"DescribeMpsUsageDataRequest": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DescribeMpsUsageData"},
		},
	},
}
