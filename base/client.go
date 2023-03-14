package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	accessKey = "STREAMLAKE_ACCESSKEY"

	defaultScheme = "https"
	defaultHost   = "vod.streamlakeapi.com"
)

var DefaultClient *http.Client

func init() {
	DefaultClient = &http.Client{
		Transport: &http.Transport{
			Proxy:           http.ProxyFromEnvironment,
			IdleConnTimeout: 10 * time.Second,
		},
	}
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}
type OpenAPI struct {
	HTTPClient  HTTPClient
	ServiceInfo ServiceInfo
	ApiInfoList map[string]*ApiInfo
}

// NewClient returns a new OpenAPI client with custom underlying http client, service info and api info map.
// If httpClient is nil, DefaultClient will be used.
func NewClient(httpClient HTTPClient, serviceInfo ServiceInfo, apiInfoMap map[string]*ApiInfo) *OpenAPI {
	if httpClient == nil {
		httpClient = DefaultClient
	}
	if serviceInfo.Scheme == "" {
		serviceInfo.Scheme = defaultScheme
	}
	if serviceInfo.Host == "" {
		serviceInfo.Host = defaultHost
	}
	if serviceInfo.Credentials.AccessKey == "" {
		// read from ENV
		if os.Getenv(accessKey) != "" {
			serviceInfo.Credentials.AccessKey = os.Getenv(accessKey)
		}
	}
	if len(apiInfoMap) == 0 {
		panic("apiInfoMap is empty")
	}

	return &OpenAPI{httpClient, serviceInfo, apiInfoMap}
}

func (c *OpenAPI) PostForAPIWithRequest(api string, reqParams interface{}) (*http.Response, error) {
	apiInfo := c.ApiInfoList[api]
	if apiInfo == nil {
		return nil, fmt.Errorf("api %s not found", api)
	}
	if apiInfo.Method != http.MethodPost {
		return nil, fmt.Errorf("api %s uses http method %s, not POST", api, apiInfo.Method)
	}

	headers := mergeHeader(c.ServiceInfo.Header, apiInfo.Header)
	ct := headers.Get("Content-Type")
	if len(ct) == 0 {
		ct = "application/x-www-form-urlencoded"
		headers.Set("Content-Type", ct)
	}
	query := apiInfo.Query
	u := url.URL{
		Scheme:   c.ServiceInfo.Scheme,
		Host:     c.ServiceInfo.Host,
		Path:     apiInfo.Path,
		RawQuery: query.Encode(),
	}
	var body io.Reader
	if ct == "application/json" {
		// set as json post in body
		if reqbytes, err := json.Marshal(reqParams); err != nil {
			return nil, err
		} else {
			body = bytes.NewReader(reqbytes)
			signatureVO := c.getSignature(api, apiInfo.Method, string(reqbytes), ct)
			signedHeaders := sign(signatureVO)
			for k, v := range signedHeaders {
				headers.Set(k, v)
			}
		}
	} else if ct == "application/x-www-form-urlencoded" {
		// set as form post in body
		if params, e := convertStructToUrlValues(reqParams); e != nil {
			return nil, e
		} else {
			body = strings.NewReader(params.Encode())
			signatureVO := c.getSignature(api, apiInfo.Method, params.Encode(), ct)
			signedHeaders := sign(signatureVO)
			for k, v := range signedHeaders {
				headers.Set(k, v)
			}
		}
	} else if len(ct) == 0 || ct == "application/octet-stream" {
		// set as query in url
		if params, e := convertStructToUrlValues(reqParams); e != nil {
			return nil, e
		} else {
			values := u.Query()
			for k, v := range params {
				for _, iv := range v {
					if values.Has(k) {
						values.Add(k, iv)
					} else {
						values.Set(k, iv)
					}
				}
			}
			u.RawQuery = values.Encode()
		}
		if ct == "application/octet-stream" {
			var err error
			body, err = getReaderFromStruct(reqParams)
			if err != nil {
				return nil, err
			}
		}
	} else {
		panic("unsupported Content-Type: " + ct)
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header = headers

	return c.HTTPClient.Do(req)
}

func (c *OpenAPI) PostForAPIWithRequestResponse(api string, req interface{}, resp interface{}) error {
	if r, e := c.PostForAPIWithRequest(api, req); e != nil {
		return fmt.Errorf("api: %s, error: %e", api, e)
	} else {
		if body, err := ioutil.ReadAll(r.Body); err != nil {
			return err
		} else {
			if r.StatusCode != http.StatusOK {
				return fmt.Errorf("api: %s, http: %d, body: %s", api, r.StatusCode, string(body))
			}
			if err2 := json.Unmarshal(body, resp); err2 != nil {
				return fmt.Errorf("api: %s, %e", api, err2)
			} else {
				return nil
			}
		}
	}
}

func (c *OpenAPI) getSignature(api string, method string, body string, contentType string) (vo SignatureVO) {
	vo = SignatureVO{}
	vo.AccessKeyId = c.ServiceInfo.Credentials.AccessKey
	vo.AccessKeySecret = c.ServiceInfo.Credentials.SecretAccessKey
	vo.Algorithm = "SL-HMAC-SHA256"
	vo.Action = api
	vo.Host = c.ServiceInfo.Host
	vo.ContentType = contentType
	vo.Region = c.ServiceInfo.Region
	vo.Version = "2022-06-23"
	vo.Service = c.ServiceInfo.ProductName
	vo.CanonicalHeaders = fmt.Sprintf("content-type:%s\nhost:%s", contentType, c.ServiceInfo.Host)
	vo.CanonicalQueryString = fmt.Sprintf("Action=%s", api)
	vo.HttpRequestMethod = method
	vo.SignedHeaders = "content-type;host"
	vo.Payload = body
	return vo

}
