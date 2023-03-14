package play

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/streamlakecloud/streamlakecloud-sdk-go/base"
)

type PlayClient struct {
	*base.OpenAPI
}

func NewPlayClient(httpClient base.HTTPClient) *PlayClient {
	return &PlayClient{
		base.NewClient(httpClient, ServiceInfo, ApiList),
	}
}

func (p *PlayClient) generateStsToken(req GenerateStsTokenRequest) (*GenerateStsTokenResponse, error) {
	resp := &GenerateStsTokenResponse{}
	err := p.PostForAPIWithRequestResponse("GenerateStsToken", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PlayClient) GeneratePlayToken(req GeneratePlayTokenRequest) (string, error) {
	ststokenresp, err := p.generateStsToken(GenerateStsTokenRequest{p.ServiceInfo.Credentials.AccessKey, req.Duration})
	if err != nil {
		return "", err
	}
	stsMap := ststokenresp.ToMap()

	urlAuthMap := req.UrlAuthInfo.ToMap()

	claims := jwt.MapClaims{
		"StsToken":    stsMap,
		"Domain":      req.Domain,
		"UrlAuthInfo": urlAuthMap,
	}
	if len(req.MediaId) > 0 {
		claims["MediaId"] = req.MediaId
	}
	// Create a new token with HS256 algorithm and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err2 := token.SignedString([]byte(req.DomainKey))
	return tokenString, err2
}
