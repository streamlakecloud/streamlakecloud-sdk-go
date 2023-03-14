package base

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type headerOptions map[string]string

func sign(vo SignatureVO) headerOptions {

	accessKeyId := vo.AccessKeyId
	secretKey := vo.AccessKeySecret
	host := vo.Host
	algorithm := vo.Algorithm
	service := vo.Service
	version := vo.Version
	action := vo.Action
	region := vo.Region
	contentType := vo.ContentType
	var timestamp = time.Now().Unix()

	// step 1: build canonical request string
	httpRequestMethod := vo.HttpRequestMethod
	canonicalURI := "/"
	canonicalQueryString := vo.CanonicalQueryString
	canonicalHeaders := vo.CanonicalHeaders
	signedHeaders := vo.SignedHeaders
	payload := vo.Payload
	hashedRequestPayload := sha256hex(payload)
	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s", httpRequestMethod, canonicalURI, canonicalQueryString, canonicalHeaders, signedHeaders, hashedRequestPayload)
	fmt.Println(canonicalRequest)

	// step 2: build string to sign
	date := time.Unix(timestamp, 0).UTC().Format("2006-01-02")
	credentialScope := fmt.Sprintf("%s/%s/sl_request", date, service)
	hashedCanonicalRequest := sha256hex(canonicalRequest)
	string2sign := fmt.Sprintf("%s\n%d\n%s\n%s", algorithm, timestamp, credentialScope, hashedCanonicalRequest)
	fmt.Println(string2sign)

	// step 3: sign string
	secretDate := hmacsha256(date, "SL"+secretKey)
	secretService := hmacsha256(service, secretDate)
	secretSigning := hmacsha256("sl_request", secretService)
	signature := hex.EncodeToString([]byte(hmacsha256(string2sign, secretSigning)))
	fmt.Println(signature)

	// step 4: build authorization
	authorization := fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%ssl_request", algorithm, accessKeyId, credentialScope, signedHeaders, signature)
	fmt.Println(authorization)

	curl := fmt.Sprintf(`curl -X POST https://%s\
 -H "Authorization: %s"\
 -H "Content-Type: application/json; charset=utf-8"\
 -H "Host: %s" -H "X-SL-Action: %s"\
 -H "X-SL-Timestamp: %d"\
 -H "X-SL-Version: %s"\
 -H "X-SL-Region: %s"\
 -d '%s'`, host, authorization, host, action, timestamp, version, region, payload)
	fmt.Println(curl)

	return headerOptions{
		"Authorization":         authorization,
		"Content-Type":          contentType,
		"Host":                  host,
		"X-SL-Action":           action,
		"X-SL-Timestamp":        strconv.FormatUint(uint64(timestamp), 10),
		"X-SL-Version":          version,
		"X-SL-Region":           region,
		"X-SL-Program-Language": "Go",
		"SignatureVersion":      "1",
		"AccessKey":             accessKeyId,
	}
}

func sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func hmacsha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return string(hashed.Sum(nil))
}
