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
	ResponseData *struct {
		TaskId string
	} `json:",omitempty"`
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

type DeleteMediaRequest struct {
	MediaId     string            `json:",omitempty"`
	DeleteItems []MediaDeleteItem `json:",omitempty"`
}
type MediaDeleteItem struct {
	Type       string `json:",omitempty"`
	TemplateId string `json:",omitempty"`
}

type DeleteMediaResponse struct {
	ResponseMeta *base.ResponseMeta
}

type DescribeMediaInfoRequest struct {
	MediaId    string   `json:",omitempty"` // 视频媒资 ID，与 PrimaryKey 二选一
	PrimaryKey string   `json:",omitempty"` // 自定义媒资 ID，与 MediaId 二选一
	Filters    []string `json:",omitempty"`
}

type ListMediaInfoRequest struct {
	MediaIds        []string `json:",omitempty"`
	CreateTimeBegin int64    `json:",omitempty"`
	CreateTimeEnd   int64    `json:",omitempty"`
	Page            int32    `json:",omitempty"`
	Size            int32    `json:",omitempty"`
}

type ListMediaInfoResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		MediaInfoList []MediaInfo
		Total         int32 `json:",omitempty"`
		Page          int32 `json:",omitempty"`
		Size          int32 `json:",omitempty"`
	} `json:",omitempty"`
}

type MediaInfo struct {
	MediaId    string
	Status     string
	Width      int32
	Height     int32
	Duration   float64
	CreateTime int64
	CoverUrl   string
	Format     string
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
type BasicInfo struct {
	SubAppId    string
	MediaId     string
	Title       string
	Description string // in format of YYYY-MM-DDThh:mm:ssZ
	CoverUrl    string
	CreateTime  string
	UpdateTime  string
}
type SourceInfo struct {
	URLPath      string
	PlayUrl      string
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
	PlayUrl             string
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
type SnapshotInfo struct {
	CoverSnapshotInfos  []CoverSnapshotInfo
	SampleSnapshotInfos []SampleSnapshotInfo
	SpriteSnapshotInfos []SpriteSnapshotInfo
	MaskSnapshotInfos   []MaskSnapshotInfo
}
type CoverSnapshotInfo struct {
	Name        string
	Type        string
	Format      string
	CdnUrl      string
	UrlPath     string
	StorageInfo StorageInfo
	Width       int32
	Height      int32
	CreateTime  int64
}
type SampleSnapshotInfo struct {
	Name                   string
	Type                   string
	Format                 string
	SampleSnapshotUrlInfos []SampleSnapshotUrlInfo
	TemplateId             string
	Width                  int32
	Height                 int32
	CreateTime             int64
}
type SampleSnapshotUrlInfo struct {
	CdnUrl      string
	UrlPath     string
	StorageInfo StorageInfo
}
type StorageInfo struct {
	StorageBucket string
	StorageKey    string
}
type SpriteSnapshotInfo struct {
	Name        string
	Type        string
	CdnUrl      string
	UrlPath     string
	TemplateId  string
	CreateTime  int64
	StorageInfo StorageInfo
}
type MaskSnapshotInfo struct {
	Name        string
	Type        string
	CdnUrl      string
	UrlPath     string
	CreateTime  string
	StorageInfo StorageInfo
}

type DescribeMediaInfoResult struct {
	MediaId        string `json:",omitempty"`
	PrimaryKey     string `json:",omitempty"`
	BasicInfo      BasicInfo
	SourceInfo     SourceInfo
	TranscodeInfos []TranscodeInfo
	SnapshotInfo   SnapshotInfo
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

type DetectMediaRequest struct {
	CallbackLink string
	MediaItemSet MediaItemSet
}

type MediaItemSet struct {
	CallbackLink string
	Bucket       string
	StoreKey     string
	ProcessTypes []string
	MediaType    string
	ClientInfo   ClientInfo
	KeyInfo      KeyInfo
	SourceUrl    string
	Text         string
}

type ClientInfo struct {
	TaskId    string
	TokenName string
	Token     string
}

type KeyInfo struct {
	EncryptionKey string
}

type DetectMediaResponse struct {
	ResponseMeta *base.ResponseMeta
}

type CreateTranscodeTemplateRequest struct {
	TranscodeTemplate TranscodeTemplate
}

type TranscodeTemplate struct {
	TemplateId           string        `json:",omitempty"`
	Name                 string        `json:",omitempty"`
	Description          string        `json:",omitempty"`
	Container            string        `json:",omitempty"`
	RemoveAudio          string        `json:",omitempty"`
	VideoTemplate        VideoTemplate `json:",omitempty"`
	AudioTemplate        AudioTemplate `json:",omitempty"`
	WatermarkTemplateIds []string      `json:",omitempty"`
}

type VideoTemplate struct {
	Codec         string `json:",omitempty"`
	Fps           int    `json:",omitempty"`
	MaxBitrate    int    `json:",omitempty"`
	LongShortMode string `json:",omitempty"`
	Width         int    `json:",omitempty"`
	Height        int    `json:",omitempty"`
	Crf           int    `json:",omitempty"`
	Gop           int    `json:",omitempty"`
}

type AudioTemplate struct {
	Codec      string `json:",omitempty"`
	Bitrate    int    `json:",omitempty"`
	SampleRate int    `json:",omitempty"`
}

type CreateTranscodeTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		TranscodeTemplateId string
	} `json:",omitempty"`
}

type UpdateTranscodeTemplateRequest struct {
	TranscodeTemplateId string            `json:",omitempty"`
	TranscodeTemplate   TranscodeTemplate `json:",omitempty"`
}

type UpdateTranscodeTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		TranscodeTemplateId string
	} `json:",omitempty"`
}

type DescribeTranscodeTemplateRequest struct {
	TranscodeTemplateId string `json:",omitempty"`
}

type DescribeTranscodeTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		CreateTime        string
		UpdateTime        string
		TranscodeTemplate TranscodeTemplate
	} `json:",omitempty"`
}

type ListTranscodeTemplateRequest struct {
	Offset int32 `json:",omitempty"`
	Limit  int32 `json:",omitempty"`
}

type ListTranscodeTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		TranscodeTemplates []TranscodeTemplate
		Limit              int32
		Offset             int32
		Total              int32
	} `json:",omitempty"`
}

type DeleteTranscodeTemplateRequest struct {
	TranscodeTemplateId string `json:",omitempty"`
}

type DeleteTranscodeTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
	} `json:",omitempty"`
}

type CreateWatermarkTemplateRequest struct {
	WatermarkTemplate WatermarkTemplate `json:",omitempty"`
}

type WatermarkTemplate struct {
	TemplateId    string        `json:",omitempty"`
	Name          string        `json:",omitempty"`
	Description   string        `json:",omitempty"`
	Type          string        `json:",omitempty"`
	ReferPosition string        `json:",omitempty"`
	MarginX       string        `json:",omitempty"`
	MarginY       string        `json:",omitempty"`
	ImageTemplate ImageTemplate `json:",omitempty"`
	TextTemplate  TextTemplate  `json:",omitempty"`
}

type ImageTemplate struct {
	Resource Resource `json:",omitempty"`
	Width    string   `json:",omitempty"`
	Height   string   `json:",omitempty"`
}

type Resource struct {
	Bucket string `json:",omitempty"`
	Object string `json:",omitempty"`
}

type TextTemplate struct {
	FontType  string `json:",omitempty"`
	Text      string `json:",omitempty"`
	FontSize  int32  `json:",omitempty"`
	FontColor string `json:",omitempty"`
}

type CreateWatermarkTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		WatermarkTemplateId string
	} `json:",omitempty"`
}

type UpdateWatermarkTemplateRequest struct {
	WatermarkTemplateId string            `json:",omitempty"`
	WatermarkTemplate   WatermarkTemplate `json:",omitempty"`
}

type UpdateWatermarkTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		WatermarkTemplateId string
	} `json:",omitempty"`
}

type DescribeWatermarkTemplateRequest struct {
	WatermarkTemplateId string `json:",omitempty"`
}

type DescribeWatermarkTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		CreateTime        string
		UpdateTime        string
		WatermarkTemplate WatermarkTemplate
	} `json:",omitempty"`
}

type ListWatermarkTemplateRequest struct {
	Offset int32 `json:",omitempty"`
	Limit  int32 `json:",omitempty"`
}

type ListWatermarkTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		WatermarkTemplates []WatermarkTemplate
		Limit              int32
		Offset             int32
		Total              int32
	} `json:",omitempty"`
}

type DeleteWatermarkTemplateRequest struct {
	WatermarkTemplateId string `json:",omitempty"`
}

type DeleteWatermarkTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
	} `json:",omitempty"`
}

type CreateSnapshotTemplateRequest struct {
	TemplateType                 string                       `json:",omitempty"`
	SnapshotByTimeOffsetTemplate SnapshotByTimeOffsetTemplate `json:",omitempty"`
	SampleSnapshotTemplate       SampleSnapshotTemplate       `json:",omitempty"`
	ImageSpriteTemplate          ImageSpriteTemplate          `json:",omitempty"`
}

type SnapshotByTimeOffsetTemplate struct {
	SnapshotTemplateId string `json:",omitempty"`
	Name               string `json:",omitempty"`
	Description        string `json:",omitempty"`
	Format             string `json:",omitempty"`
	Width              int32  `json:",omitempty"`
	Height             int32  `json:",omitempty"`
	OffsetTime         int32  `json:",omitempty"`
}

type SampleSnapshotTemplate struct {
	SnapshotTemplateId string `json:",omitempty"`
	Name               string `json:",omitempty"`
	Description        string `json:",omitempty"`
	SampleType         string `json:",omitempty"`
	Interval           int32  `json:",omitempty"`
	Format             string `json:",omitempty"`
	Count              int64  `json:",omitempty"`
	Width              int32  `json:",omitempty"`
	Height             int32  `json:",omitempty"`
	OffsetTime         int32  `json:",omitempty"`
}

type ImageSpriteTemplate struct {
	SnapshotTemplateId string `json:",omitempty"`
	Name               string `json:",omitempty"`
	Description        string `json:",omitempty"`
	SampleType         string `json:",omitempty"`
	SampleInterval     int32  `json:",omitempty"`
	RowCount           int32  `json:",omitempty"`
	ColumnCount        int32  `json:",omitempty"`
	Width              int32  `json:",omitempty"`
	Height             int32  `json:",omitempty"`
	Format             string `json:",omitempty"`
}

type CreateSnapshotTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		SnapshotTemplateId string
	} `json:",omitempty"`
}

type UpdateSnapshotTemplateRequest struct {
	TemplateType                 string                       `json:",omitempty"`
	SnapshotTemplateId           string                       `json:",omitempty"`
	SnapshotByTimeOffsetTemplate SnapshotByTimeOffsetTemplate `json:",omitempty"`
	SampleSnapshotTemplate       SampleSnapshotTemplate       `json:",omitempty"`
	ImageSpriteTemplate          ImageSpriteTemplate          `json:",omitempty"`
}

type UpdateSnapshotTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		SnapshotTemplateId string
		TemplateType       string
	} `json:",omitempty"`
}

type DescribeSnapshotTemplateRequest struct {
	TemplateType       string `json:",omitempty"`
	SnapshotTemplateId string `json:",omitempty"`
}

type DescribeSnapshotTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		CreateTime                   string
		UpdateTime                   string
		TemplateType                 string
		SnapshotByTimeOffsetTemplate SnapshotByTimeOffsetTemplate
		SampleSnapshotTemplate       SampleSnapshotTemplate
		ImageSpriteTemplate          ImageSpriteTemplate
	} `json:",omitempty"`
}

type ListSnapshotTemplateRequest struct {
	TemplateType string `json:",omitempty"`
	Offset       int32  `json:",omitempty"`
	Limit        int32  `json:",omitempty"`
}

type ListSnapshotTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		SnapshotByTimeOffsetTemplates []SnapshotByTimeOffsetTemplate
		SampleSnapshotTemplates       []SampleSnapshotTemplate
		ImageSpriteTemplates          []ImageSpriteTemplate
	} `json:",omitempty"`
}

type DeleteSnapshotTemplateRequest struct {
	TemplateType       string `json:",omitempty"`
	SnapshotTemplateId string `json:",omitempty"`
}

type DeleteSnapshotTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
	} `json:",omitempty"`
}

type CreateWorkflowTemplateRequest struct {
	MediaProcessWorkflowTemplate MediaProcessWorkflowTemplate `json:",omitempty"`
}

type MediaProcessWorkflowTemplate struct {
	WorkflowId                string                     `json:",omitempty"`
	WorkflowName              string                     `json:",omitempty"`
	Description               string                     `json:",omitempty"`
	TranscodeTasks            []TranscodeTask            `json:",omitempty"`
	SnapshotByTimeOffsetTasks []SnapshotByTimeOffsetTask `json:",omitempty"`
	SampleSnapshotTasks       []SampleSnapshotTask       `json:",omitempty"`
	ImageSpriteTasks          []ImageSpriteTask          `json:",omitempty"`
}

type TranscodeTask struct {
	TemplateId string `json:",omitempty"`
}

type SnapshotByTimeOffsetTask struct {
	TemplateId string `json:",omitempty"`
}

type SampleSnapshotTask struct {
	TemplateId string `json:",omitempty"`
}

type ImageSpriteTask struct {
	TemplateId string `json:",omitempty"`
}

type CreateWorkflowTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		WorkflowId string
	} `json:",omitempty"`
}

type UpdateWorkflowTemplateRequest struct {
	WorkflowId                   string                       `json:",omitempty"`
	MediaProcessWorkflowTemplate MediaProcessWorkflowTemplate `json:",omitempty"`
}

type UpdateWorkflowTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		WorkflowId string
	} `json:",omitempty"`
}

type ListWorkflowTemplateRequest struct {
	Names  []string `json:",omitempty"`
	Offset int32    `json:",omitempty"`
	Limit  int32    `json:",omitempty"`
}

type ListWorkflowTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		MediaProcessWorkflowTemplates []MediaProcessWorkflowTemplate
		Limit                         int32
		Offset                        int32
		Total                         int32
	} `json:",omitempty"`
}

type DeleteWorkflowTemplateRequest struct {
	WorkflowId string `json:",omitempty"`
}

type DeleteWorkflowTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
	} `json:",omitempty"`
}

type DescribeTaskDetailRequest struct {
	TaskId string `json:",omitempty"`
}

type DescribeTaskDetailResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		TaskType           string
		CreateTime         string
		BeginProcessTime   string
		FinishTime         string
		WorkflowTaskResult WorkflowTaskResult
	} `json:",omitempty"`
}

type WorkflowTaskResult struct {
	TaskId                  string                   `json:",omitempty"`
	Status                  string                   `json:",omitempty"`
	MediaId                 string                   `json:",omitempty"`
	Metadata                Metadata                 `json:",omitempty"`
	MediaProcessTaskResults []MediaProcessTaskResult `json:",omitempty"`
}

type Metadata struct {
	FileSize        int64         `json:",omitempty"`
	Height          int32         `json:",omitempty"`
	Width           int32         `json:",omitempty"`
	Bitrate         int32         `json:",omitempty"`
	Duration        float64       `json:",omitempty"`
	Fps             int32         `json:",omitempty"`
	Format          string        `json:",omitempty"`
	VideoMaxBitrate int64         `json:",omitempty"`
	VideoStreams    []VideoStream `json:",omitempty"`
	AudioStreams    []AudioStream `json:",omitempty"`
}

type MediaProcessTaskResult struct {
	Type                           string                         `json:",omitempty"`
	TranscodeTaskResult            TranscodeTaskResult            `json:",omitempty"`
	SnapshotByTimeOffsetTaskResult SnapshotByTimeOffsetTaskResult `json:",omitempty"`
	SampleSnapshotTaskResult       SampleSnapshotTaskResult       `json:",omitempty"`
	ImageSpriteTaskResult          ImageSpriteTaskResult          `json:",omitempty"`
}

type TranscodeTaskResult struct {
	Status           string             `json:",omitempty"`
	Input            TranscodeTaskInput `json:",omitempty"`
	BeginProcessTime string             `json:",omitempty"`
	FinishTime       string             `json:",omitempty"`
}

type SnapshotByTimeOffsetTaskResult struct {
	Status           string                        `json:",omitempty"`
	Input            SnapshotByTimeOffsetTaskInput `json:",omitempty"`
	BeginProcessTime string                        `json:",omitempty"`
	FinishTime       string                        `json:",omitempty"`
}

type SampleSnapshotTaskResult struct {
	Status           string                  `json:",omitempty"`
	Input            SampleSnapshotTaskInput `json:",omitempty"`
	BeginProcessTime string                  `json:",omitempty"`
	FinishTime       string                  `json:",omitempty"`
}

type ImageSpriteTaskResult struct {
	Status           string               `json:",omitempty"`
	Input            ImageSpriteTaskInput `json:",omitempty"`
	BeginProcessTime string               `json:",omitempty"`
	FinishTime       string               `json:",omitempty"`
}

type TranscodeTaskInput struct {
	TemplateId string `json:",omitempty"`
}

type SnapshotByTimeOffsetTaskInput struct {
	TemplateId string `json:",omitempty"`
}

type SampleSnapshotTaskInput struct {
	TemplateId string `json:",omitempty"`
}

type ImageSpriteTaskInput struct {
	TemplateId string `json:",omitempty"`
}

type ApplyUploadInfoRequest struct {
	SessionKey string `json:",omitempty"`
	MediaSort  string `json:",omitempty"`
	FilePath   string `json:",omitempty"`
	Format     string `json:",omitempty"`
}

type ApplyUploadInfoResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		SessionKey    string
		UploadAddress UploadAddress
		UploadAuth    UploadAuth
	} `json:",omitempty"`
}

type UploadAddress struct {
	StorageBucket  string `json:",omitempty"`
	Region         string `json:",omitempty"`
	UploadEndpoint string `json:",omitempty"`
	UploadPath     string `json:",omitempty"`
}

type UploadAuth struct {
	SecretId    string `json:",omitempty"`
	SecretKey   string `json:",omitempty"`
	Token       string `json:",omitempty"`
	ExpiredTime int64  `json:",omitempty"`
}

type CommitUploadRequest struct {
	SessionKey string `json:",omitempty"`
}

type CommitUploadResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		MediaId   string `json:",omitempty"`
		MediaSort string `json:",omitempty"`
	} `json:",omitempty"`
}
