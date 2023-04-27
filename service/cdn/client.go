package cdn

import "github.com/streamlakecloud/streamlakecloud-sdk-go/base"

type CDNClient struct {
	*base.OpenAPI
}

func NewCDNClient(httpClient base.HTTPClient) *CDNClient {
	return &CDNClient{
		base.NewClient(httpClient, ServiceInfo, ApiList),
	}
}

func NewCDNClientV2(httpClient base.HTTPClient, serviceInfo base.ServiceInfo) *CDNClient {
	return &CDNClient{
		base.NewClient(httpClient, serviceInfo, ApiList),
	}
}

func (c *CDNClient) DescribeCdnLogs(req DescribeCdnLogsRequest) (*DescribeCdnLogsResponse, error) {
	resp := &DescribeCdnLogsResponse{}
	err := c.OpenAPI.PostForAPIWithRequestResponse("DescribeCdnLogs", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CDNClient) PreloadObjectCaches(req PreloadObjectCachesRequest) (*PreloadObjectCachesResponse, error) {
	resp := &PreloadObjectCachesResponse{}
	err := c.OpenAPI.PostForAPIWithRequestResponse("PreloadObjectCaches", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CDNClient) RefreshObjectCaches(req RefreshObjectCachesRequest) (*RefreshObjectCachesResponse, error) {
	resp := &RefreshObjectCachesResponse{}
	err := c.OpenAPI.PostForAPIWithRequestResponse("RefreshObjectCaches", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CDNClient) DescribeRefreshTasks(req DescribeRefreshTasksRequest) (*DescribeRefreshTasksResponse, error) {
	resp := &DescribeRefreshTasksResponse{}
	err := c.OpenAPI.PostForAPIWithRequestResponse("DescribeRefreshTasks", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CDNClient) GetDomainRealTimeCdnData(req GetDomainRealTimeCdnRequest) (*GetDomainRealTimeCdnResponse, error) {
	resp := &GetDomainRealTimeCdnResponse{}
	err := c.OpenAPI.PostForAPIWithRequestResponse("DescribeDomainRealTimeCdnData", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CDNClient) GetDomainRealTimeOriginData(req GetDomainRealTimeOriginRequest) (*GetDomainRealTimeOriginResponse, error) {
	resp := &GetDomainRealTimeOriginResponse{}
	err := c.OpenAPI.PostForAPIWithRequestResponse("DescribeDomainRealTimeOriginData", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CDNClient) PushPCDNCache(req PushPCDNObjectCacheRequest) (*PushPCDNObjectCacheResponse, error) {
	resp := &PushPCDNObjectCacheResponse{}
	err := c.OpenAPI.PostForAPIWithRequestResponse("PushPCDNObjectCache", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CDNClient) ListPcdnDataSources(req ListPcdnDataSourcesRequest) (*ListPcdnDataSourcesResponse, error) {
	resp := &ListPcdnDataSourcesResponse{}
	err := c.OpenAPI.PostForAPIWithRequestResponse("ListPcdnDataSources", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CDNClient) DescribePcdnDataSummary(req DescribePcdnDataSummaryRequest) (*DescribePcdnDataSummaryResponse, error) {
	resp := &DescribePcdnDataSummaryResponse{}
	err := c.OpenAPI.PostForAPIWithRequestResponse("DescribePcdnDataSummary", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CDNClient) DescribePcdnDataDetail(req DescribePcdnDataDetailRequest) (*DescribePcdnDataDetailResponse, error) {
	resp := &DescribePcdnDataDetailResponse{}
	err := c.OpenAPI.PostForAPIWithRequestResponse("DescribePcdnDataDetail", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
