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
	CallbackURL       string `json:"CallbackURL,omitempty"`
}
type FetchUploadRequest struct {
	URLSets   []FetchUploadURLSet
	SpaceName string `json:"SpaceName,omitempty"`
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
	SpaceName      string `json:"SpaceName,omitempty"`
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
	SpaceName     string `json:"SpaceName,omitempty"`
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

// deprecated, do not use
type DescribeMediaProcessJobsRequest struct {
	JobIds string
}

type DescribeMediaProcessJobRequest struct {
	JobId       string
	ProcessType string `json:"ProcessType,omitempty"` // "MediaFeatureAnalysis"
	SpaceName   string `json:"SpaceName,omitempty"`
}

type TranscodeJobInfo struct {
	MediaId string
	TranscodeInfo
}
type CustomMediaProcessJobInfo struct {
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

// deprecated, do not use
type DescribeMediaProcessJobsResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		JobInfos       []CustomMediaProcessJobInfo
		NonExistJobIds []string
	} `json:",omitempty"`
}

type MediaProcessVideoStreamInfo struct {
	CodeType     string  // "video" 文件类型
	CodecName    string  // "hevc" 编码器名称
	Width        int     // 1080
	Height       int     // 1920
	AvgFrameRate string  // "24135680/804527" 平均帧率
	RFrameRate   string  // "30/1" 帧率
	Bitrate      int     // 12800 码率
	StartTime    float32 // 0 开始时间
	PixFmt       string  // "yuv420p" PixelFormat 像素格式
	HdrType      string  // "SDR" hdr类型
	Duration     float32 // 314.268359 时长 单位:秒
}

type MediaProcessAudioStreamInfo struct {
	CodeType   string  // "audio" 文件类型
	StartTime  float32 // 0 开始时间
	Duration   float32 // 314.268359 时长 单位:秒
	CodecName  string  // "aac" 编码器名称
	Channels   int     // 2 声道数
	SampleRate int     // 48000 采样率
	Bitrate    int     // 12800 码率
}

type MediaProcessFormat struct {
	FormatName string  // "MP4" 格式名称
	ListSize   int     // m3u8文件的大小，单位: byte
	Size       int     // 文件大小，单位: byte
	Duration   float32 // 总时长，单位秒
	Bitrate    int     // 869880 总码率，单位: bps
}

type MediaProcessMediaInfo struct {
	UrlPath      string
	VideoStreams []MediaProcessVideoStreamInfo
	AudioStreams []MediaProcessAudioStreamInfo
	Format       MediaProcessFormat
	ErrorCode    string
	ErrorMessage string
}

type QualityFeature struct {
	NrQualityAvgScore     float32   // 无参考质量评价分数，取值区间：[1-5]
	BlurProbability       float32   // 模糊概率，取值区间：[0-1]
	BlockyProbability     float32   // 块效应程度，取值区间：[0-1]
	NoiseProbability      float32   // 噪声程度，取值区间：[0-1]
	InterlaceRatio        float32   // 交错视频概率，取值区间：[0-1]
	DirtylensProbability  float32   // 镜头脏污程度，取值区间：[0-1]
	BlurProbabilityNew    float32   // 新版本模型的模糊概率，越高表示越模糊，取值区间：[0-1]
	BlockyProbabilityNew  float32   // 新版本模型的有块效应的概率，越高表示块效应程度越大，取值区间：[0-1]
	LowLightPercentage    float32   // 暗光的概率，越高表示越可能是暗光情况，取值区间：[0-1]。视频专属的特征值。
	CorruptionProbability float32   // 花屏的概率，越高表示越可能是花屏，取值区间：[0-1]
	NoiseProbabilityNew   float32   // 新版本模型的噪声概率，越高表示越可能有噪声，取值区间：[0-1]
	InterlaceScaleClass   float32   // interlace比例，取值区间：[0-1]
	NrQualityScoreNN      []float32 // 无参考质量评价分数，取值区间：[1-5]
}

type AestheticsFeature struct {
	DefocusProbability      float32 // 大光圈概率，取值区间：[0-1]
	Colorfulness            float32 // 色彩丰富度，取值区间：[0-1]
	UnderexposedProbability float32 // 曝光不足程度，取值区间：[0-1]
	OverexposedProbability  float32 // 曝光过度程度，取值区间：[0-1]
	LowContrastProbability  float32 // 低对比度程度，取值区间：[0-1]
	AestheticsScore         float32 // 美学评分，取值区间：[0-1]
}

type ContentFeature struct {
	JunkVideoProbability float32   // 内容无意义，取值区间：[0-1]
	LetterboxProbability float32   // 三段式视频，取值区间：[0-1]
	OritationProbability float32   // 横屏视频概率，取值区间：[0-1]
	R0Probability        float32   // 横屏的概率，取值区间：[0-1]
	R90Probability       float32   // 旋转90度的概率，取值区间：[0-1]
	R270Probability      float32   // 旋转270度的概率，取值区间：[0-1]
	NormalProbability    float32   // 普通画面概率，取值区间：[0-1]
	PureProbability      float32   // 纯色画面概率，取值区间：[0-1]
	TextProbability      float32   // 文字画面概率，取值区间：[0-1]
	CartoonProbability   float32   // 卡通画面概率，取值区间：[0-1]
	QrcodeProbability    float32   // 二维码画面概率，取值区间：[0-1]
	LetterboxTop         []float32 // 画面上边占比，取值区间：[0-1]
	LetterboxBottom      []float32 // 画面下边占比，取值区间：[0-1]
	FaceRatio            float32   // 视频中含有人脸的视频帧占比，取值区间：[0-1]
	VrVideo              int       // 1表示为vr视频，0表示不是
}

type AudioFeature struct {
	AudioQuality               float32 // 音频质量，取值区间：[0-1]
	MusicProbability           float32 // 音频为音乐概率，取值区间：[0-1]
	DialogProbability          float32 // 音频为对话概率，取值区间：[0-1]
	BackgroundSoundProbability float32 // 音频为背景音概率，取值区间：[0-1]
	MutePercentage             float32 // 无声占比，取值区间：[0-1]
}

type CodingFeature struct {
	EntropyAvg         float32   // 帧内复杂度的均值，越大表示图像/视频帧复杂度越高，取值范围大致在-6到26。
	EntropyVar         float32   // 帧内复杂度的方差，越大表示视频内复杂度变化范围越大，取值 : >0。
	IntraComplexity    float32   // 帧内复杂度，取值范围：>0
	InterComplexity    float32   // 帧间复杂度，取值范围：>0
	TopIntraComplexity float32   // 多帧帧内复杂度TOP部分，取值范围：>0
	TopInterComplexity float32   // 多帧帧间复杂度TOP部分，取值范围：>0
	MaxIntraComplexity float32   // 多帧帧内复杂度最大值，取值范围：>0
	MaxInterComplexity float32   // 多帧帧间复杂度最大值，取值范围：>0
	MinIntraComplexity float32   // 多帧帧内复杂度最小值，取值范围：>0
	MinInterComplexity float32   // 多帧帧间复杂度最小值，取值范围：>0
	YMean              float32   // Y通道均值，取值范围：>0，视频专属的特征值。
	YMeanMax           float32   // 多帧Y通道均值最大值，取值范围：>0，视频专属的特征值。
	YMeanMin           float32   // 多帧Y通道均值最小值，取值范围：>=0，视频专属的特征值。
	YDiffStd           float32   // Y通道差异标准差，取值范围：>0，视频专属的特征值。
	YDiffStdMax        float32   // 多帧Y通道差异标准差最大值，取值范围：>0，视频专属的特征值。
	YDiffStdMin        float32   // 多帧Y通道差异标准差最小值，取值范围：>=0，视频专属的特征值。
	CapeMotionVertical float32   // 垂直运动量，取值范围：>0
	CapeMotionHorizon  float32   // 水平运动量，取值范围：>0
	Embedding          []float32 // 编码特征向量，无范围限制
}

type SceneFeature struct {
	OthersProbability    float32 // 场景-其他类别概率，取值区间：[0-1]
	SkyProbability       float32 // 场景-天空画面概率，取值区间：[0-1]
	MountainProbability  float32 // 场景-高山画面概率，取值区间：[0-1]
	FieldProbability     float32 // 场景-田野画面概率，取值区间：[0-1]
	WoodProbability      float32 // 场景-森林画面概率，取值区间：[0-1]
	WaterProbability     float32 // 场景-水画面概率，取值区间：[0-1]
	CrowdProbability     float32 // 场景-人群画面概率，取值区间：[0-1]
	PersonProbability    float32 // 场景-全身人像概率，取值区间：[0-1]
	PortraitProbability  float32 // 场景-半身人像概率，取值区间：[0-1]
	CarProbability       float32 // 场景-汽车概率，取值区间：[0-1]
	WriteProbability     float32 // 场景-文本概率，取值区间：[0-1]
	FoodProbability      float32 // 场景-美食画面概率，取值区间：[0-1]
	CgProbability        float32 // 场景-cg画面概率，取值区间：[0-1]
	NightProbability     float32 // 场景-夜景画面概率，取值区间：[0-1]
	PurecolorProbability float32 // 场景-纯色图片概率，取值区间：[0-1]
	CapProbability       float32 // 场景-字幕概率，取值区间：[0-1]
	MobaProbability      float32 // 场景-王者荣耀概率，取值区间：[0-1]
	FpsProbability       float32 // 场景-吃鸡游戏概率，取值区间：[0-1]
	StageProbability     float32 // 场景-舞台类别概率，取值区间：[0-1]
}

type MediaProcessMediaFeature struct {
	QualityFeature    QualityFeature    // 画质特征
	AestheticsFeature AestheticsFeature // 美学特征
	ContentFeature    ContentFeature    // 内容特征
	AudioFeature      AudioFeature      // 音频特征
	CodingFeature     CodingFeature     // 编码特征
	SceneFeature      SceneFeature      // 场景特征
	Version           string            // "premium"，特征版本
}

type MediaProcessJobInfo struct {
	JobId        string
	TemplateId   string
	UrlPaths     string // comma separated
	MediaInfos   []MediaProcessMediaInfo
	MediaFeature MediaProcessMediaFeature
	ErrorCode    string
	ErrorMessage string
}
type DescribeMediaProcessJobResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		CreateTime  string // in format of YYYY-MM-DDThh:mm:ssZ
		ProcessTime string // in format of YYYY-MM-DDThh:mm:ssZ
		FinishTime  string // in format of YYYY-MM-DDThh:mm:ssZ
		ProcessType string
		Status      string // one of "pending", "processing", "completed", "error"
		JobInfo     MediaProcessJobInfo
	} `json:",omitempty"`
}

type ListMediaProcessJobRequest struct {
	ProcessType string `json:",omitempty"`          // "MediaFeatureAnalysis"
	Offset      int32  `json:",omitempty"`          // 分页偏移量，默认值：0。不能小于0
	Limit       int32  `json:",omitempty"`          // 分页返回条数，最大不超过50。不能小于0，默认值：20
	SpaceName   string `json:"SpaceName,omitempty"` // 应用空间，默认为：default_space
}

type ShortJobInfo struct {
	JobId       string  // 任务 ID
	KvqScore    float32 // KVQ总体得分
	UrlPaths    string  // 转出文件的url或者目录，多个逗号(,)分割
	ProcessType string
	Version     string // "basic" or "premium"。KVQ任务版本
	CreateTime  string // in format of YYYY-MM-DDThh:mm:ssZ
	ProcessTime string // in format of YYYY-MM-DDThh:mm:ssZ
	FinishTime  string // in format of YYYY-MM-DDThh:mm:ssZ
	Status      string // 任务状态。可选值：pending：排队中、processing: 处理中、completed: 执行成功、error: 执行失败
}

type ListMediaProcessJobResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		ShortJobInfos []ShortJobInfo // 对应的简短任务结果列表
		Limit         int32          // 返回的模板个数
		Offset        int32          // 下次分页起始坐标
		Total         int32          // 满足查询条件的模板总数
	} `json:",omitempty"`
}

/******************************************************************************
 * Manage *
 *****************************************************************************/

type DeleteMediaRequest struct {
	MediaId     string            `json:",omitempty"`
	DeleteItems []MediaDeleteItem `json:",omitempty"`
	SpaceName   string            `json:",omitempty"`
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
	SpaceName  string   `json:",omitempty"`
}

type ListMediaInfoRequest struct {
	MediaIds        []string `json:",omitempty"`
	CreateTimeBegin int64    `json:",omitempty"`
	CreateTimeEnd   int64    `json:",omitempty"`
	Page            int32    `json:",omitempty"`
	Size            int32    `json:",omitempty"`
	SpaceName       string   `json:",omitempty"`
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

type VideoStreamOutput struct {
	Duration float64
	Width    int32
	Height   int32
	Fps      int32
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
	SpaceName  string `json:"SpaceName,omitempty"`
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
	SpaceName         string `json:"SpaceName,omitempty"`
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
	SpaceName           string            `json:"SpaceName,omitempty"`
}

type UpdateTranscodeTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		TranscodeTemplateId string
	} `json:",omitempty"`
}

type DescribeTranscodeTemplateRequest struct {
	TranscodeTemplateId string `json:",omitempty"`
	SpaceName           string `json:"SpaceName,omitempty"`
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
	Offset    int32  `json:",omitempty"`
	Limit     int32  `json:",omitempty"`
	SpaceName string `json:"SpaceName,omitempty"`
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
	SpaceName           string `json:"SpaceName,omitempty"`
}

type DeleteTranscodeTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
	} `json:",omitempty"`
}

type CreateWatermarkTemplateRequest struct {
	WatermarkTemplate WatermarkTemplate `json:",omitempty"`
	SpaceName         string            `json:"SpaceName,omitempty"`
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
	SpaceName           string            `json:"SpaceName,omitempty"`
}

type UpdateWatermarkTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		WatermarkTemplateId string
	} `json:",omitempty"`
}

type DescribeWatermarkTemplateRequest struct {
	WatermarkTemplateId string `json:",omitempty"`
	SpaceName           string `json:"SpaceName,omitempty"`
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
	Offset    int32  `json:",omitempty"`
	Limit     int32  `json:",omitempty"`
	SpaceName string `json:"SpaceName,omitempty"`
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
	SpaceName           string `json:"SpaceName,omitempty"`
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
	SpaceName                    string                       `json:"SpaceName,omitempty"`
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
	SpaceName                    string                       `json:"SpaceName,omitempty"`
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
	SpaceName          string `json:"SpaceName,omitempty"`
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
	SpaceName    string `json:"SpaceName,omitempty"`
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
	SpaceName          string `json:"SpaceName,omitempty"`
}

type DeleteSnapshotTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
	} `json:",omitempty"`
}

type CreateWorkflowTemplateRequest struct {
	MediaProcessWorkflowTemplate MediaProcessWorkflowTemplate `json:",omitempty"`
	SpaceName                    string                       `json:"SpaceName,omitempty"`
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
	SpaceName                    string                       `json:"SpaceName,omitempty"`
}

type UpdateWorkflowTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		WorkflowId string
	} `json:",omitempty"`
}

type ListWorkflowTemplateRequest struct {
	Names     []string `json:",omitempty"`
	Offset    int32    `json:",omitempty"`
	Limit     int32    `json:",omitempty"`
	SpaceName string   `json:"SpaceName,omitempty"`
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
	SpaceName  string `json:"SpaceName,omitempty"`
}

type DeleteWorkflowTemplateResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
	} `json:",omitempty"`
}

type DescribeTaskDetailRequest struct {
	TaskId    string `json:",omitempty"`
	SpaceName string `json:"SpaceName,omitempty"`
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
	FileSize        int64               `json:",omitempty"`
	Height          int32               `json:",omitempty"`
	Width           int32               `json:",omitempty"`
	Bitrate         int32               `json:",omitempty"`
	Duration        float64             `json:",omitempty"`
	Fps             int32               `json:",omitempty"`
	Format          string              `json:",omitempty"`
	VideoMaxBitrate int64               `json:",omitempty"`
	VideoStreams    []VideoStreamOutput `json:",omitempty"`
	AudioStreams    []AudioStream       `json:",omitempty"`
}

type MediaProcessTaskResult struct {
	Type                           string                         `json:",omitempty"`
	TranscodeTaskResult            TranscodeTaskResult            `json:",omitempty"`
	SnapshotByTimeOffsetTaskResult SnapshotByTimeOffsetTaskResult `json:",omitempty"`
	SampleSnapshotTaskResult       SampleSnapshotTaskResult       `json:",omitempty"`
	ImageSpriteTaskResult          ImageSpriteTaskResult          `json:",omitempty"`
}

type TranscodeTaskResult struct {
	Status           string              `json:",omitempty"`
	Input            TranscodeTaskInput  `json:",omitempty"`
	Output           TranscodeTaskOutput `json:",omitempty"`
	BeginProcessTime string              `json:",omitempty"`
	FinishTime       string              `json:",omitempty"`
}

type SnapshotByTimeOffsetTaskResult struct {
	Status           string                         `json:",omitempty"`
	Input            SnapshotByTimeOffsetTaskInput  `json:",omitempty"`
	Output           SnapshotByTimeOffsetTaskOutput `json:",omitempty"`
	BeginProcessTime string                         `json:",omitempty"`
	FinishTime       string                         `json:",omitempty"`
}

type SampleSnapshotTaskResult struct {
	Status           string                   `json:",omitempty"`
	Input            SampleSnapshotTaskInput  `json:",omitempty"`
	Output           SampleSnapshotTaskOutput `json:",omitempty"`
	BeginProcessTime string                   `json:",omitempty"`
	FinishTime       string                   `json:",omitempty"`
}

type ImageSpriteTaskResult struct {
	Status           string                `json:",omitempty"`
	Output           ImageSpriteTaskOutput `json:",omitempty"`
	Input            ImageSpriteTaskInput  `json:",omitempty"`
	BeginProcessTime string                `json:",omitempty"`
	FinishTime       string                `json:",omitempty"`
}

type TranscodeTaskInput struct {
	TemplateId string `json:",omitempty"`
}

type TranscodeTaskOutput struct {
	URLPath      string              `json:",omitempty"`
	FileSize     int64               `json:",omitempty"`
	Height       int32               `json:",omitempty"`
	Width        int32               `json:",omitempty"`
	Bitrate      int32               `json:":,omitempty"`
	Duration     float32             `json:",omitempty"`
	Fps          int32               `json:",omitempty"`
	Format       string              `json:",omitempty"`
	VideoStreams []VideoStreamOutput `json:",omitempty"`
	AudioStreams []AudioStream       `json:",omitempty"`
}

type SnapshotByTimeOffsetTaskInput struct {
	TemplateId string `json:",omitempty"`
}

type SnapshotByTimeOffsetTaskOutput struct {
	URLPath  string `json:",omitempty"`
	FileSize int64  `json:",omitempty"`
	Height   int32  `json:",omitempty"`
	Width    int32  `json:",omitempty"`
	Format   string `json:",omitempty"`
}

type SampleSnapshotTaskInput struct {
	TemplateId string `json:",omitempty"`
}

type SampleSnapshotTaskOutput struct {
	SampleSnapshotInfos []SampleSnapshotTaskInfo `json:",omitempty"`
}

type SampleSnapshotTaskInfo struct {
	URLPath   string  `json:",omitempty"`
	TimeStamp float64 `json:",omitempty"`
	Width     int32   `json:",omitempty"`
	Height    int32   `json:",omitempty"`
}

type ImageSpriteTaskInput struct {
	TemplateId string `json:",omitempty"`
}

type ImageSpriteTaskOutput struct {
	URLPath string `json:",omitempty"`
}

type ApplyUploadInfoRequest struct {
	SessionKey string `json:",omitempty"`
	MediaSort  string `json:",omitempty"`
	FilePath   string `json:",omitempty"`
	Format     string `json:",omitempty"`
	SpaceName  string `json:",omitempty"`
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

type ListDanmakuPreAuditRequest struct {
	StartTime  int64  `json:"StartTime,omitempty"`
	EndTime    int64  `json:"EndTime,omitempty"`
	ResourceId string `json:"ResourceId,omitempty"`
	PrimaryKey string `json:"PrimaryKey,omitempty"`
	Limit      int32  `json:"Limit,omitempty"`
}

type ListDanmakuPreAuditResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		List  []DanmakuInfo `json:"List"`
		Count int32         `json:"Count"`
		Total int32         `json:"Total"`
	}
}

type DanmakuInfo struct {
	DanmakuId  int64  `json:"DanmakuId"`
	ResourceId string `json:"ResourceId"`
	UserId     string `json:"UserId"`
	Body       string `json:"Body"`
	Position   int64  `json:"Position"`
	Color      string `json:"Color"`
	Size       int32  `json:"Size"`
}

type UpdateDanmakuAuditResultRequest struct {
	DanmakuId  int64   `json:"DanmakuId,omitempty"`
	DanmakuIds []int64 `json:"DanmakuIds,omitempty"`
	Status     string  `json:"Status,omitempty"`
}

type UpdateDanmakuAuditResultResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *struct {
		DanmakuId  int64   `json:"DanmakuId"`
		DanmakuIds []int64 `json:"DanmakuIds"`
	}
}
