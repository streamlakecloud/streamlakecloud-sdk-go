package metric

import "github.com/streamlakecloud/streamlakecloud-sdk-go/base"

type DescribeDomainUsageDataRequest struct {
	DomainName string // 加速域名。若参数为空，默认返回所有加速域名合并后数据。支持批量查询(去重)，多个用半角逗号（,）分隔。(主域名)
	StartTime  string // 获取数据起始时间点。格式为：yyyy-MM-ddTHH:mm:ssZ（UTC时间）。
	EndTime    string // 获取数据结束时间点，需晚于起始时间。格式为：yyyy-MM-ddTHH:mm:ssZ（UTC时间）。
	Internal   string // 统计时间粒度。取值：5minutes：5分钟 (最大时间范围1天,最长过去7天) hour：小时。(最大时间范围7天,最长过去30天) day：天。(最大时间范围30天, 最长过去90天)默认按时间跨度决定，小于等于1小时以5分钟为粒度，小于等于1天以1小时为粒度，大于1天则以天为粒度。
	Percentile string // 查询95峰值, 可选值：DomesticBps, OverseaBps, TotalBps, 返回值在返回记录UsageDataItem对应字段中
}

type DescribeDomainUsageDataResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData interface{} `json:",omitempty"`
}

type DescribeMpsUsageDataRequest struct {
	StartTime   string // 获取数据起始时间点。格式为：yyyy-MM-ddTHH:mm:ssZ（UTC时间）。
	EndTime     string // 获取数据结束时间点，需晚于起始时间。格式为：yyyy-MM-ddTHH:mm:ssZ（UTC时间）。
	Region      string // 存储地域。默认返回所有区域。支持批量查询，多个地域使用半角逗号（,）分隔。取值：华北地区(北京): cn-beijing
	Internal    string // 统计时间粒度。取值：hour：小时。(最大时间范围7天,最长过去30天) day：天。(最大时间范围30天, 最长过去90天)默认按时间跨度决定，小于等于1天以1小时为粒度，大于1天则以天为粒度。
	ProcessType string // 媒体处理的分类: Frame: 截图; Transcode: 转码; AudioEnhance: 音频增强; AGC: 响度均衡; SDR2HDR: SDR转HDR
}

type DescribeMpsUsageDataResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData interface{} `json:",omitempty"`
}

type DescribeStorageUsageDataRequest struct {
	StartTime string // 获取数据起始时间点。格式为：yyyy-MM-ddTHH:mm:ssZ（UTC时间）。
	EndTime   string // 获取数据结束时间点，需晚于起始时间。格式为：yyyy-MM-ddTHH:mm:ssZ（UTC时间）。
	Region    string // 存储地域。默认返回所有地域。支持批量查询，多个区域用半角逗号（,）分隔。取值：华北地区(北京): cn-beijing 目前只要华北
	Bucket    string // 存储桶的名字,暂时还不支持
	Interval  string // 统计时间粒度。取值：5minutes：5分钟 (最大时间范围1天,最长过去7天) hour：小时。(最大时间范围7天,最长过去30天) day：天。(最大时间范围30天, 最长过去90天) 默认按时间跨度决定，小于等于1小时以5分钟为粒度，小于等于1天以1小时为粒度，大于1天则以天为粒度。
}

type DescribeStorageUsageDataResponse struct {
	ResponseMeta *base.ResponseMeta
	ResponseData interface{} `json:",omitempty"`
}
