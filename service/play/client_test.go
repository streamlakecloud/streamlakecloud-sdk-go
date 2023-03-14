package play

import (
	"os"
	"testing"
)

var (
	DEMO_HOST_STAGING      = os.Getenv("STREAMLAKE_VOD_GO_SDK_ENDPOINT")
	DEMO_TEST_ACCESS_KEY   = os.Getenv("STREAMLAKE_VOD_GO_SDK_ACCESS_KEY2")
	DEMO_TEST_SECRET_KEY   = os.Getenv("STREAMLAKE_VOD_GO_SDK_SECRET_KEY")
	DEMO_TEST_ACCESS_TOKEN = os.Getenv("STREAMLAKE_VOD_GO_SDK_ACCESS_TOKEN")
)

func TestGenerateStsToken(t *testing.T) {
	ply := NewPlayClient(nil)
	ply.ServiceInfo.Host = DEMO_HOST_STAGING
	ply.ServiceInfo.Credentials.AccessKey = DEMO_TEST_ACCESS_KEY
	ply.ServiceInfo.Credentials.SecretAccessKey = DEMO_TEST_SECRET_KEY
	ply.ServiceInfo.Credentials.AccessToken = DEMO_TEST_ACCESS_TOKEN
	req := GenerateStsTokenRequest{
		ply.ServiceInfo.Credentials.AccessKey,
		10000}

	if resp, err := ply.generateStsToken(req); err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("%+v", resp)
	}
}

func TestGeneratePlayToken(t *testing.T) {
	ply := NewPlayClient(nil)
	ply.ServiceInfo.Host = DEMO_HOST_STAGING
	ply.ServiceInfo.Credentials.AccessKey = DEMO_TEST_ACCESS_KEY
	ply.ServiceInfo.Credentials.SecretAccessKey = DEMO_TEST_SECRET_KEY
	ply.ServiceInfo.Credentials.AccessToken = DEMO_TEST_ACCESS_TOKEN

	req := GeneratePlayTokenRequest{
		Domain:    "test.com",
		DomainKey: "test",
		MediaId:   "12412frwaftwet",
		Duration:  300,
	}
	if token, err := ply.GeneratePlayToken(req); err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("token %+v", token)
	}
}
