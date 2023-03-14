package cdn

import "github.com/streamlakecloud/streamlakecloud-sdk-go/base"

type DescribeCdnLogsRequest struct {
	StartTime  string // in format of "yyyy-MM-dd HH:mm:ss"
	EndTime    string // in format of "yyyy-MM-dd HH:mm:ss"
	DomainName string // "example.com"
}

type DescribeCdnLogsResult struct {
	CDNLogs []string
}
type DescribeCdnLogsResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *DescribeCdnLogsResult `json:",omitempty"`
}

type PreloadObjectCachesRequest struct {
	ObjectPath string
}

type PreloadObjectCachesResult struct {
	PreloadTaskId string
}

type PreloadObjectCachesResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *PreloadObjectCachesResult `json:",omitempty"`
}

type RefreshObjectCachesRequest struct {
	ObjectPath string
	ObjectType string
}

type RefreshObjectCachesResult struct {
	RefreshTaskId string
}

type RefreshObjectCachesResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *RefreshObjectCachesResult `json:",omitempty"`
}

type DescribeRefreshTasksRequest struct {
	TaskId   string
	TaskType string
	Offset   int32
	Limit    int32
}

type Tasks struct {
	CreationTime string
	ObjectPath   string
	ObjectType   string
	Status       string
	TaskId       string
	TaskType     string
}

type DescribeRefreshTasksResult struct {
	Offset     int32
	Limit      int32
	TotalCount int32
	Tasks      []Tasks
}

type DescribeRefreshTasksResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *DescribeRefreshTasksResult `json:",omitempty"`
}

type GetDomainRealTimeCdnRequest struct {
	DomainName string
	StartTime  string
	EndTime    string
	Metric     string
	Interval   string
}

type RealTimeDataItem struct {
	TimeStamp string
	Value     interface{}
}

type GetDomainRealTimeCdnResult struct {
	DomainName string
	Interval   string
	Metric     string
	Data       []RealTimeDataItem
}

type GetDomainRealTimeCdnResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *GetDomainRealTimeCdnResult
}

type GetDomainRealTimeOriginRequest struct {
	DomainName string
	StartTime  string
	EndTime    string
	Metric     string
	Interval   string
}

type GetDomainRealTimeOriginResult struct {
	DomainName string
	Interval   string
	Metric     string
	Data       []RealTimeDataItem
}

type GetDomainRealTimeOriginResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData *GetDomainRealTimeOriginResult
}

type PushPCDNObjectCacheRequest struct {
	FileSets []PCDNFileSet
}

type PCDNFileSet struct {
	URLPath    string
	FileSize   int64
	Crc32      string
	BlockCrc32 []string
	Qps        int64
}

type PushPCDNObjectCacheResponse struct {
	ResponseMeta *base.ResponseMeta
}
