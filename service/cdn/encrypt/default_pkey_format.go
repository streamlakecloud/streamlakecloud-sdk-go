package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"hash/crc32"
	"math/rand"
	"net"
	"net/url"
	"strings"
	"time"
)

const (
	AES_KEY_SIZE = 32
	AES_IV_SIZE  = 16
	IPV4_FLAG    = 0x00
	IPV6_FLAG    = 0x01
	DEFAULT_TTL  = 6 * 60 * 60
)

var (
	DEFAULT_IP         = make([]byte, 16)
	BYTES_12           = make([]byte, 12)
	VERSION            = []byte{0x00, 0x05}
	DEFAULT_SERVICE_ID = []byte{0x00, 0x00}
)

type DefaultPkeyFormat struct{}

func NewDefaultPkeyFormat() *DefaultPkeyFormat {
	return &DefaultPkeyFormat{}
}

func (b DefaultPkeyFormat) format(request *PkeyRequest) (string, error) {
	if request == nil {
		return "", errors.New("pkey request must not be null")
	}

	pkeyData, err := pkeyData(request)
	if err != nil {
		return "", err
	}

	encBytes, err := encryptAES256(request.GetCryptoKey(), pkeyData)
	if err != nil {
		return "", err
	}

	versionBytes := version()
	concatenatedBytes := append(versionBytes, encBytes...)

	encodedString := base64.RawURLEncoding.WithPadding(base64.NoPadding).EncodeToString(concatenatedBytes)
	return encodedString, nil
}

func pkeyData(request *PkeyRequest) ([]byte, error) {
	var outputStream bytes.Buffer
	outputStream.Write(intToByteArray(rand.Int()))
	outputStream.Write(intToByteArray(int(currentTimeSecond())))
	outputStream.Write(intToByteArray(ttl(request.ttl)))
	if request.serviceId > 0 {
		outputStream.Write(shortToByteArray(request.serviceId))
	} else {
		outputStream.Write(DEFAULT_SERVICE_ID)
	}
	outputStream.Write([]byte{0})
	outputStream.Write(intToByteArray(pathCrc32(request.uri)))
	outputStream.Write(shortToByteArray(request.limitSpeeds))
	outputStream.Write(shortToByteArray(request.limitTimeSeconds))
	outputStream.Write(shortToByteArray(request.limitTSIdx))
	outputStream.Write(shortToByteArray(int16(len(request.gid))))
	outputStream.Write([]byte(request.gid))
	writeIpAddress(&outputStream, request.ipAddress)
	outputBytes := outputStream.Bytes()
	crc32AsInt := crc32Helper(outputBytes, 0, len(outputBytes))
	outputStream.Write(intToByteArray(crc32AsInt))
	return outputStream.Bytes(), nil
}

func version() []byte {
	return VERSION
}

func intToByteArray(value int) []byte {
	return []byte{
		byte(value >> 24),
		byte(value >> 16),
		byte(value >> 8),
		byte(value),
	}
}

func shortToByteArray(value int16) []byte {
	return []byte{
		byte(value >> 8),
		byte(value),
	}
}

func currentTimeSecond() int64 {
	return time.Now().Unix()
}

func ttl(ttl int) int {
	if ttl > 0 {
		return ttl
	} else {
		return DEFAULT_TTL
	}
}

func pathCrc32(uri *url.URL) int {
	path := uri.Path
	var outputBytes []byte
	if strings.Contains(path, ".") {
		urlWithoutSuffix := path[0:strings.Index(path, ".")]
		outputBytes = []byte(urlWithoutSuffix)
	} else {
		outputBytes = []byte(path)
	}
	return crc32Helper(outputBytes, 0, len(outputBytes))
}

func crc32Helper(value []byte, offset, length int) int {
	crc := crc32.NewIEEE()
	crc.Write(value[offset : offset+length])
	return int(crc.Sum32())
}

func writeIpAddress(outputStream *bytes.Buffer, ipAddress net.IP) {
	if ipAddress == nil {
		outputStream.WriteByte(IPV4_FLAG)
		outputStream.Write(DEFAULT_IP)
	} else {
		if ipAddress.To4() != nil {
			outputStream.WriteByte(IPV4_FLAG)
			outputStream.Write(ipAddress.To4())
			outputStream.Write(BYTES_12)
		} else {
			outputStream.WriteByte(IPV6_FLAG)
			outputStream.Write(ipAddress.To16())
		}
	}
}

func encryptAES256(cryptoKey *CryptoKey, plainText []byte) ([]byte, error) {
	aesKey, err := hex.DecodeString(cryptoKey.aesKey)
	if err != nil || len(aesKey) != AES_KEY_SIZE {
		return nil, fmt.Errorf("invalid aes key size")
	}
	aesIv, err := hex.DecodeString(cryptoKey.aesIv)
	if err != nil || len(aesIv) != AES_IV_SIZE {
		return nil, fmt.Errorf("invalid aes iv size")
	}
	return aesCBC256Encrypt(aesKey, aesIv, plainText)
}

func aesCBC256Encrypt(key, iv, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	paddedData := addPadding(data, aes.BlockSize)

	ciphertext := make([]byte, len(paddedData))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedData)

	return ciphertext, nil
}

func addPadding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	paddingData := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, paddingData...)
}
