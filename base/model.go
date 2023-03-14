package base

import (
	"net/http"
	"net/url"
)

// ResponseMeta is a typical response from API server
type ResponseMeta struct {
	RequestId    string `json:",omitempty"`
	ErrorCode    string `json:",omitempty"`
	ErrorMessage string `json:",omitempty"`
}

type Credentials struct {
	AccessKey       string
	SecretAccessKey string
	Service         string
	Region          string
	SessionToken    string
	HttpClient      HTTPClient
	AccessToken     string
}

type ServiceInfo struct {
	Region      string      // not used for now.
	Scheme      string      // http or https
	Host        string      // The OpenAPI host.
	Header      http.Header // service headers
	Credentials Credentials // service credentials
	ProductName string      // product name. eg. vod
}

type ApiInfo struct {
	Method string      // http.MethodGet, http.MethodPost, etc
	Path   string      // default corresponding path
	Query  url.Values  // Api default Query values to be merged
	Form   url.Values  // Api default Form values to be merged
	Header http.Header // Api default Header values to be merged
}

type SignatureVO struct {
	AccessKeyId          string // StreamLake密钥ak
	AccessKeySecret      string // StreamLake密钥sk
	Algorithm            string // StreamLake加签算法
	Service              string // StreamLake 服务编码
	Host                 string // StreamLake request host
	ContentType          string // StreamLake request content-type
	Region               string // StreamLake request region
	Action               string // StreamLake request action
	Version              string // StreamLake request version
	HttpRequestMethod    string // HTTP 请求方法（GET、POST ）
	CanonicalQueryString string // 发起 HTTP 请求 URL 中的查询字符串，
	CanonicalHeaders     string // 参与签名的头部信息
	SignedHeaders        string // 参与签名的头部信息，说明此次请求有哪些头部参与了签名
	Payload              string // 请求正文
}
