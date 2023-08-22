package encrypt

import (
	"net"
	"net/url"
)

type PkeyRequest struct {
	uri              *url.URL
	cryptoKey        *CryptoKey
	serviceId        int16
	platformId       int16
	gid              string
	ipAddress        net.IP
	ttl              int
	limitSpeeds      int16
	limitTimeSeconds int16
	limitTSIdx       int16
}

func NewPkeyRequest(builder *PkeyRequestBuilder) *PkeyRequest {
	return &PkeyRequest{
		uri:              builder.uri,
		cryptoKey:        builder.cryptoKey,
		serviceId:        builder.serviceId,
		platformId:       builder.platformId,
		gid:              builder.gid,
		ipAddress:        builder.ipAddress,
		ttl:              builder.ttl,
		limitSpeeds:      builder.limitSpeeds,
		limitTimeSeconds: builder.limitTimeSeconds,
		limitTSIdx:       builder.limitTSIdx,
	}
}

func (p *PkeyRequest) GetUri() *url.URL {
	return p.uri
}

func (p *PkeyRequest) GetCryptoKey() *CryptoKey {
	return p.cryptoKey
}

func (p *PkeyRequest) GetServiceId() int16 {
	return p.serviceId
}

func (p *PkeyRequest) GetPlatformId() int16 {
	return p.platformId
}

func (p *PkeyRequest) GetGid() string {
	return p.gid
}

func (p *PkeyRequest) GetIpAddress() net.IP {
	return p.ipAddress
}

func (p *PkeyRequest) GetTtl() int {
	return p.ttl
}

func (p *PkeyRequest) GetLimitSpeeds() int16 {
	return p.limitSpeeds
}

func (p *PkeyRequest) GetLimitTimeSeconds() int16 {
	return p.limitTimeSeconds
}

func (p *PkeyRequest) GetLimitTSIdx() int16 {
	return p.limitTSIdx
}

type CryptoKey struct {
	aesKey string
	aesIv  string
}

func NewCryptoKey(aesKey, aesIv string) *CryptoKey {
	return &CryptoKey{
		aesKey: aesKey,
		aesIv:  aesIv,
	}
}

func (c *CryptoKey) GetAesKey() string {
	return c.aesKey
}

func (c *CryptoKey) GetAesIv() string {
	return c.aesIv
}

type PkeyRequestBuilder struct {
	uri              *url.URL
	cryptoKey        *CryptoKey
	serviceId        int16
	platformId       int16
	gid              string
	ipAddress        net.IP
	ttl              int
	limitSpeeds      int16
	limitTimeSeconds int16
	limitTSIdx       int16
}

func NewPkeyRequestBuilder() *PkeyRequestBuilder {
	return &PkeyRequestBuilder{}
}

func (b *PkeyRequestBuilder) WithUri(uri *url.URL) *PkeyRequestBuilder {
	b.uri = uri
	return b
}

func (b *PkeyRequestBuilder) WithCryptoKey(cryptoKey *CryptoKey) *PkeyRequestBuilder {
	b.cryptoKey = cryptoKey
	return b
}

func (b *PkeyRequestBuilder) WithServiceId(serviceId int16) *PkeyRequestBuilder {
	b.serviceId = serviceId
	return b
}

func (b *PkeyRequestBuilder) WithPlatformId(platformId int16) *PkeyRequestBuilder {
	b.platformId = platformId
	return b
}

func (b *PkeyRequestBuilder) WithGid(gid string) *PkeyRequestBuilder {
	b.gid = gid
	return b
}

func (b *PkeyRequestBuilder) WithIpAddress(ipAddress net.IP) *PkeyRequestBuilder {
	b.ipAddress = ipAddress
	return b
}

func (b *PkeyRequestBuilder) WithTtl(ttl int) *PkeyRequestBuilder {
	b.ttl = ttl
	return b
}

func (b *PkeyRequestBuilder) WithLimitSpeeds(limitSpeeds int16) *PkeyRequestBuilder {
	b.limitSpeeds = limitSpeeds
	return b
}

func (b *PkeyRequestBuilder) WithLimitTimeSeconds(limitTimeSeconds int16) *PkeyRequestBuilder {
	b.limitTimeSeconds = limitTimeSeconds
	return b
}

func (b *PkeyRequestBuilder) WithLimitTSIdx(limitTSIdx int16) *PkeyRequestBuilder {
	b.limitTSIdx = limitTSIdx
	return b
}

func (b *PkeyRequestBuilder) Build() *PkeyRequest {
	return NewPkeyRequest(b)
}
