package encrypt

import (
	"fmt"
	"testing"
)

func TestEncryptCdnUrl(t *testing.T) {
	key := "xxx"
	iv := "xxx"
	url := "xxx"
	client := NewEncryptClient(key, iv)
	res, _ := client.SignCdnUrl(url, *NewExtraInfoBuilder().SetClientIP("127.0.0.1").SetTTL(600).Build())
	fmt.Println(string(res))
}
