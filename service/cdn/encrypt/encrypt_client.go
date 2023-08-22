package encrypt

import (
	"net"
	"net/url"
)

type EncryptClient struct {
	Key string
	Iv  string
}

func NewEncryptClient(key, iv string) *EncryptClient {
	return &EncryptClient{
		Key: key,
		Iv:  iv,
	}
}

func (client *EncryptClient) SignCdnUrl(originalUrl string, extraInfo ExtraInfo) (string, error) {
	request := PkeyRequest{
		uri:       urlMustParse(originalUrl),
		cryptoKey: &CryptoKey{aesKey: client.Key, aesIv: client.Iv},
		ttl:       extraInfo.TTL,
		ipAddress: net.ParseIP(extraInfo.ClientIP),
	}

	parameters := buildPkeyParameter(&request)
	pkey := parameters["pkey"]

	return originalUrl + "?pkey=" + pkey, nil
}

func (client *EncryptClient) SetKey(key string) {
	client.Key = key
}

func (client *EncryptClient) SetIv(iv string) {
	client.Iv = iv
}

func urlMustParse(rawURL string) *url.URL {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return parsedURL
}
