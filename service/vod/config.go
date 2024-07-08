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
	"ListMediaInfo": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"ListMediaInfo"},
		},
	},
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
	"DeleteMedia": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DeleteMedia"},
		},
	},

	/*************
	 * 媒资上传
	 *************/
	"ApplyUploadInfo": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"ApplyUploadInfo"},
		},
	},
	"CommitUpload": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"CommitUpload"},
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
	"DetectMedia": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DetectMedia"},
		},
	},
	/*************
	 * 模板和任务流
	 *************/
	"CreateTranscodeTemplate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"CreateTranscodeTemplate"},
		},
	},
	"UpdateTranscodeTemplate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"UpdateTranscodeTemplate"},
		},
	},
	"DescribeTranscodeTemplate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DescribeTranscodeTemplate"},
		},
	},
	"ListTranscodeTemplate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"ListTranscodeTemplate"},
		},
	},
	"DeleteTranscodeTemplate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DeleteTranscodeTemplate"},
		},
	},
	"CreateWatermarkTemplate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"CreateWatermarkTemplate"},
		},
	},
	"UpdateWatermarkTemplate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"UpdateWatermarkTemplate"},
		},
	},
	"DescribeWatermarkTemplate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DescribeWatermarkTemplate"},
		},
	},
	"ListWatermarkTemplate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"ListWatermarkTemplate"},
		},
	},
	"DeleteWatermarkTemplate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DeleteWatermarkTemplate"},
		},
	},
	"CreateSnapshotTemplate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"CreateSnapshotTemplate"},
		},
	},
	"UpdateSnapshotTemplate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"UpdateSnapshotTemplate"},
		},
	},
	"DescribeSnapshotTemplate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DescribeSnapshotTemplate"},
		},
	},
	"ListSnapshotTemplate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"ListSnapshotTemplate"},
		},
	},
	"DeleteSnapshotTemplate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DeleteSnapshotTemplate"},
		},
	},
	"CreateWorkflowTemplate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"CreateWorkflowTemplate"},
		},
	},
	"UpdateWorkflowTemplate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"UpdateWorkflowTemplate"},
		},
	},
	"ListWorkflowTemplate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"ListWorkflowTemplate"},
		},
	},
	"DeleteWorkflowTemplate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DeleteWorkflowTemplate"},
		},
	},
	"DescribeTaskDetail": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"DescribeTaskDetail"},
		},
	},
	/*************
	 * 弹幕审核
	 *************/
	"ListDanmakuPreAudit": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"ListDanmakuPreAudit"},
		},
	},
	"UpdateDanamkuAuditResult": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action": []string{"UpdateDanamkuAuditResult"},
		},
	},
}
