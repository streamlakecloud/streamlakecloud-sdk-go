package metric

import "github.com/streamlakecloud/streamlakecloud-sdk-go/base"

type MetricClient struct {
	*base.OpenAPI
}

func NewMetricClient(httpClient base.HTTPClient) *MetricClient {
	return &MetricClient{
		base.NewClient(httpClient, ServiceInfo, ApiList),
	}
}
func (v *MetricClient) DescribeDomainUsageData(req DescribeDomainUsageDataRequest) (*DescribeDomainUsageDataResponse, error) {
	resp := &DescribeDomainUsageDataResponse{}
	err := v.PostForAPIWithRequestResponse("DescribeDomainUsageData", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *MetricClient) DescribeMpsUsageData(req DescribeMpsUsageDataRequest) (*DescribeMpsUsageDataResponse, error) {
	resp := &DescribeMpsUsageDataResponse{}
	err := v.PostForAPIWithRequestResponse("DescribeMpsUsageData", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *MetricClient) DescribeStorageUsageData(req DescribeStorageUsageDataRequest) (*DescribeStorageUsageDataResponse, error) {
	resp := &DescribeStorageUsageDataResponse{}
	err := v.PostForAPIWithRequestResponse("DescribeStorageUsageData", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
