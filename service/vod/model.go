package vod

import "github.com/streamlakecloud/streamlakecloud-sdk-go/base"

/******************************************************************************
 * Upload *
 *****************************************************************************/

type FetchUploadURLSet struct {
	SourceURL         string
	MediaType         string `json:",omitempty"`
	AttachedMediaType string `json:",omitempty"`
	AttachedKey       string `json:",omitempty"`
	CallbackArgs      string `json:",omitempty"`
	WorkflowId        string `json:",omitempty"`
	PrimaryKey        string `json:",omitempty"`
}
type FetchUploadRequest struct {
	URLSets []FetchUploadURLSet
}

type FetchUploadResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData interface{} `json:",omitempty"`
}

type FetchStoreURLSet struct {
	SourceURL    string
	StoreKey     string
	StoreBucket  string `json:",omitempty"`
	CallbackArgs string `json:",omitempty"`
	CallbackURL  string `json:",omitempty"`
}
type FetchStoreRequest struct {
	URLSets []FetchStoreURLSet
}

type FetchStoreJobInfo struct {
	JobId     string
	SourceURL string
}

type FetchStoreResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		JobInfos []FetchStoreJobInfo
	} `json:",omitempty"`
}

type DescribeFetchJobsRequest struct {
	JobIds []string
}

type DescribeFetchJobsJobInfo struct {
	JobId  string
	Status string // one of "RUNNING", "COMPLETED", "FAILED", "TIMED_OUT", "TERMINATED", "PAUSED", "UNRECOGNIZED", "QUERY_FAILED"
}
type DescribeFetchJobsResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		JobInfos []DescribeFetchJobsJobInfo
	} `json:",omitempty"`
}

type ApplyUploadTokenRequest struct {
	MediaType string
}

type ApplyUploadTokenResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		UploadEndpoint  string
		UploadSignature string
	} `json:",omitempty"`
}

type VerifyUploadTokenRequest struct {
	VodUploadToken string
}

type VerifyUploadTokenResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		MediaId string
	} `json:",omitempty"`
}

/******************************************************************************
 * Process *
 *****************************************************************************/

type SubmitWorkflowRequest struct {
	MediaId        string `json:",omitempty"`
	PrimaryKey     string `json:",omitempty"`
	WorkflowId     string
	OverrideParams string `json:",omitempty"`
	CallbackArgs   string `json:",omitempty"`
}

type SubmitWorkflowResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData interface{} `json:",omitempty"`
}

type WatermarkSet struct {
	WatermarkTemplateId string `json:",omitempty"`
	ResourceKey         string `json:",omitempty"`
	MarginX             string `json:",omitempty"` // 0
	MarginY             string `json:",omitempty"`
	Width               string `json:",omitempty"`
	Height              string `json:",omitempty"`
	ReferPosition       string `json:",omitempty"` // "topRight"
	WidthReferEdge      string `json:",omitempty"`
	HeightReferEdge     string `json:",omitempty"`
	MarginXReferEdge    string `json:",omitempty"`
	MarginYReferEdge    string `json:",omitempty"`
}
type TranscodeSet struct {
	TranscodeTemplateId string         //"480P_MP4_H265_0"
	URLPath             string         // "/SD/d91ds03r-31234-t1111.mp4"
	Container           string         `json:",omitempty"` // "mp4"
	Codec               string         `json:",omitempty"` // "h264" or "h265"
	Fps                 string         `json:",omitempty"` // 25
	Width               string         `json:",omitempty"` // 1080
	Height              string         `json:",omitempty"` // 1920
	Gop                 string         `json:",omitempty"` // 250
	MaxRate             string         `json:",omitempty"`
	LongShortMode       string         `json:",omitempty"` // true
	WatermarkSets       []WatermarkSet `json:",omitempty"`
}
type ProcessMediaRequest struct {
	MediaId       string
	PrimaryKey    string
	CallbackArgs  string
	TranscodeSets []TranscodeSet
}

type ProcessMediaJobInfo struct {
	JobId               string
	TranscodeTemplateId string
}
type ProcessMediaResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		JobInfos       []ProcessMediaJobInfo
		NonExistJobIds []string
	} `json:",omitempty"`
}

type DescribeMediaProcessJobsRequest struct {
	JobIds string
}

type TranscodeJobInfo struct {
	MediaId string
	TranscodeInfo
}
type MediaProcessJobInfo struct {
	JobId            string
	JobType          string
	CreateTime       string
	FinishTime       string
	Status           string // one of "SUBMITTED", "PROCESSING", "SUCCESS", "FAILED", "CANCELED"
	ErrorCode        string
	ErrorMessage     string
	Progress         int32            // value of percentage, 0~100
	TranscodeJobInfo TranscodeJobInfo // only meaningful when SUCCESS
}
type DescribeMediaProcessJobsResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		JobInfos       []MediaProcessJobInfo
		NonExistJobIds []string
	} `json:",omitempty"`
}

/******************************************************************************
 * Manage *
 *****************************************************************************/

type DescribeMediaInfoRequest struct {
	MediaId    string // 视频媒资 ID，与 PrimaryKey 二选一
	PrimaryKey string // 自定义媒资 ID，与 MediaId 二选一
}

type AudioStream struct {
	Duration   float64
	Bitrate    int32
	Channels   int32
	SampleRate int32
	Codec      string
}
type VideoStream struct {
	Duration float64
	Width    int32
	Height   int32
	Fps      float32
	Bitrate  int32
	Rotate   string // one of {"0", "90", "180", "270"}
	Codec    string
}
type SourceInfo struct {
	URLPath      string
	Format       string
	Duration     float64
	CreateTime   string // in format of YYYY-MM-DDThh:mm:ssZ
	Width        int32
	Height       int32
	Fps          float32
	Bitrate      int32
	FileSize     int64
	VideoStreams []VideoStream
	AudioStreams []AudioStream
	HdrType      string // enum, one of {"SDR", "HDR10", "HDR10+", "Dolby Vision", "HLG", "SDR+"}
}

type TranscodeInfo struct {
	TranscodeTemplateId string
	URLPath             string
	Format              string
	Duration            float64
	CreateTime          string // in format of YYYY-MM-DDThh:mm:ssZ
	Width               int32
	Height              int32
	Fps                 float32
	FileSize            int64
	VideoMaxBitrate     int32 // 峰值码率，单位：Kbps
	VideoStreams        []VideoStream
	AudioStreams        []AudioStream
	HdrType             string // enum, one of {"SDR", "HDR10", "HDR10+", "Dolby Vision", "HLG", "SDR+"}
	Bitrate             int32  // 文件码率，单位：Kbps
}

type DescribeMediaInfoResult struct {
	MediaId        string `json:",omitempty"`
	PrimaryKey     string `json:",omitempty"`
	SourceInfo     SourceInfo
	TranscodeInfos []TranscodeInfo
}

type DescribeMediaInfoResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData DescribeMediaInfoResult `json:",omitempty"`
}

type DescribeAttachedMediaInfoRequest struct {
	MediaKeys string // string. list of comma separated MediaKey
}

type InputFileSet struct {
	Bucket string
	Object string
	Url    string
}

type OutputFileSet struct {
	Bucket string
	Object string
}

type OperationSet struct {
	TemplateId    string
	ProcessType   string
	InputFileSet  InputFileSet
	OutputFileSet OutputFileSet
	ExtraParams   map[string]string
}

type ProcessSet struct {
	OperationSets  []OperationSet
	CallbackUrl    string
	CallbackMethod string
	UserData       string
}

type SubmitMediaProcessJobsRequest struct {
	MediaId    string
	ProcessSet ProcessSet
}

type SubmitMediaProcessJobsJobInfo struct {
	JobId      string
	TemplateId string
}

type SubmitMediaProcessJobsResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		JobInfos []SubmitMediaProcessJobsJobInfo
	} `json:",omitempty"`
}

type AttachedMediaInfo struct {
	AttachedMediaId   string
	AttachedMediaType string
	AttachedMediaKey  string
}
type DescribeAttachedMediaInfoResult struct {
	AttachedMedias []AttachedMediaInfo
	NonExistedKeys []string
}

type DescribeAttachedMediaInfoResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData DescribeAttachedMediaInfoResult `json:",omitempty"`
}

type DescribePlayQualityDataSourcesRequest struct {
	StartTime   string
	EndTime     string
	Metric      string
	QueryFilter []string
}

type DescribePlayQualityDataSourcesResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		Metric string
		Filter PlayQualityFilterResult
	} `json:",omitempty"`
}

type PlayQualityFilterResult struct {
	Province   []string
	ISP        []string
	Network    []string
	Platform   []string
	AppVersion []string
	Codec      []string
	Resolution []string
}

type DescribePlayQualitySummaryRequest struct {
	StartTime string
	EndTime   string
	Filters   PlayQualityFilterInfo
	Metric    string
}

type PlayQualityFilterInfo struct {
	Domains    []string
	Province   []string
	ISP        []string
	Network    []string
	Platform   []string
	AppVersion []string
	Codec      []string
	Resolution []string
}

type DescribePlayQualitySummaryResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		Metric string
		Data   DescribePlayQualityOverViewDataInfo
	} `json:",omitempty"`
}

type DescribePlayQualityOverViewDataInfo struct {
	PlayPerformance DescribePlayQualityOverViewPlayPerformanceInfo
	UserExperience  DescribePlayQualityOverViewUserExperienceInfo
	CdnDownLoad     DescribePlayQualityOverViewCdnDownLoadInfo
}

type DescribePlayQualityOverViewPlayPerformanceInfo struct {
	PlayCount                   string
	ExperienceFirstScreen       string
	PlayerFirstScreen           string
	StartPlayFailedRate         string
	VSF                         string
	EBVS                        string
	FrameLossRate               string
	FrameLossHundredSeconds     string
	AvgBitrate                  string
	BlockRate                   string
	BlockTimeHundredSeconds     string
	BlockDurationHundredSeconds string
	FaultAfterPlayRate          string
}

type DescribePlayQualityOverViewUserExperienceInfo struct {
	DeviceNum             string
	AvgPlayNumByDevice    string
	TotalPlayDuration     string
	AvgPlayDuration       string
	AvgPlayDurationDevice string
	CompleteRate          string
	PlayCompleteRate      string
}

type DescribePlayQualityOverViewCdnDownLoadInfo struct {
	CdnFinishedRate  string
	CdnFailedRate    string
	CdnCancelledRate string
	CdnSlowedRate    string
}

type DescribePlayQualityDetailRequest struct {
	StartTime string
	EndTime   string
	Filters   PlayQualityFilterInfo
	Metric    string
	Interval  string
	Dimension []string
	Top       string
	Sort      string
}

type DescribePlayQualityDetailResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		Metric string
		Data   []DescribePlayQualityDetailDataInfo
	} `json:",omitempty"`
}

type DescribePlayQualityDetailDataInfo struct {
	DimensionValue string
	ValueItem      []DescribePlayQualityDetailDataValueItem
}

type DescribePlayQualityDetailDataValueItem struct {
	Value     string
	TimeStamp string
}
