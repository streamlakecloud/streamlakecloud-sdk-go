package play

import (
	"net/http"
	"net/url"
	"os"

	"github.com/streamlakecloud/streamlakecloud-sdk-go/base"
)

var ServiceInfo = base.ServiceInfo{
	Region: "cn-beijing",
	Scheme: "https",
	Host:   "vod.streamlakeapi.com",
	Header: http.Header{
		"Content-Type": []string{"application/json"},
		"Accept":       []string{"application/json"},
	},
	Credentials: base.Credentials{},
}

var playAccessKey = os.Getenv("STREAMLAKE_VOD_GO_SDK_PLAY_ACCESS_KEY")
var ApiList = map[string]*base.ApiInfo{
	// Play
	"GenerateStsToken": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"GenerateStsToken"},
		},
		Header: http.Header{
			"AccessKey": []string{playAccessKey}, // to be override from Credentials.AccessToken
		},
	},
}
