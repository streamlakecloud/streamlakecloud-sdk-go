package vod

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
	/*************
	 * 媒体上传
	 *************/
	"DescribeFetchJobs": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DescribeFetchJobs"},
		},
	},

	"FetchStore": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"FetchStore"},
		},
	},

	"FetchUpload": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"FetchUpload"},
		},
	},

	"ApplyUploadToken": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"ApplyUploadToken"},
		},
	},

	"VerifyUploadToken": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"VerifyUploadToken"},
		},
	},

	/*************
	 * 媒体处理
	 *************/
	"SubmitWorkflow": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"SubmitWorkflow"},
		},
	},

	"ProcessMedia": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"ProcessMedia"},
		},
	},

	"SubmitMediaProcessJobs": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"SubmitMediaProcessJobs"},
		},
	},

	"DescribeMediaProcessJobs": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DescribeMediaProcessJobs"},
		},
	},
	/*************
	 * 媒资管理
	 *************/
	"DescribeMediaInfo": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DescribeMediaInfo"},
		},
	},
	"DescribeAttachedMediaInfo": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DescribeAttachedMediaInfo"},
		},
	},
	/*************
	 * 播放质量
	 *************/
	"DescribePlayQualityDataSources": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DescribePlayQualityDataSources"},
		},
	},
	"DescribePlayQualitySummary": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DescribePlayQualitySummary"},
		},
	},
	"DescribePlayQualityDetail": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DescribePlayQualityDetail"},
		},
	},
}
