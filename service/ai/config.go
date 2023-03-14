package ai

import (
	"net/http"
	"net/url"

	"github.com/streamlakecloud/streamlakecloud-sdk-go/base"
)

var ServiceInfo = base.ServiceInfo{
	Region:      "cn-beijing",
	Host:        "vod.streamlakeapi.com",
	Scheme:      "https",
	Credentials: base.Credentials{},
}

var ApiList = map[string]*base.ApiInfo{
	// AI
	"ComposeVideo": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"ComposeVideo"},
		},
	},
	"ScanImage": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"ScanImage"},
		},
	},
}
