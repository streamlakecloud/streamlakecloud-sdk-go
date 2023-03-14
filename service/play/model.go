package play

import (
	"time"
)

type GenerateStsTokenRequest struct {
	AccessKeyId string `json:"accessKeyId,omitempty"`
	Duration    uint64 `json:"duration,omitempty"`
}

type GenerateStsTokenResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  *struct {
		AccessKeyId     string `json:"accessKeyId,omitempty"`
		AccessKeySecret string `json:"aAccessKeySecret,omitempty"`
		SessionToken    string `json:"sessionToken,omitempty"`
		Expiration      int64  `json:"expiration,omitempty"`
	} `json:"result,omitempty"`
}

func (r GenerateStsTokenResponse) ToMap() map[string]interface{} {
	if r.Result == nil {
		return map[string]interface{}{}
	}
	m := map[string]interface{}{
		"TempAccessKey": r.Result.AccessKeyId,
		"TempSecretKey": r.Result.AccessKeySecret,
		"SecurityToken": r.Result.SessionToken,
		"CreateTime":    time.Now().Unix(),
		"ExpireTime":    r.Result.Expiration / 1000,
		"StsTokenVer":   2,
	}
	return m
}

type UrlAuthInfo struct {
	ExpireTime int64
}

func (u *UrlAuthInfo) ToMap() map[string]int64 {
	m := make(map[string]int64)
	if u.ExpireTime > 0 {
		m["ExpireTime"] = u.ExpireTime
	}
	return m
}

type GeneratePlayTokenRequest struct {
	Domain      string
	DomainKey   string
	MediaId     string
	Duration    uint64
	UrlAuthInfo UrlAuthInfo
}

type GeneratePlayTokenResponse string
