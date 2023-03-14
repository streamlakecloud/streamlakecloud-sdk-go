package ai

import (
	"os"
	"testing"
)

var (
	DEMO_HOST_STAGING    = os.Getenv("STREAMLAKE_VOD_GO_SDK_ENDPOINT")
	DEMO_TEST_ACCESS_KEY = os.Getenv("STREAMLAKE_VOD_GO_SDK_ACCESS_KEY")
)

var TEMPLATES = [...]string{"artLine",
	"cherry",
	"cute",
	"cuteThings",
	"dynamicVideo",
	"focus",
	"geometry",
	"memoryFilm",
	"memoryTime",
	"millennialGirl",
	"player",
	"summer",
	"timeImprint",
	"travelDiary",
	"vigor",
}

var IMAGES = [...]string{"group1_face_50_30.jpg",
	"group1_face_360p.png",
	"480X720.jpeg",
	"group3_face_540p.jpg",
	"group2_noface_720p.png",
	"face_special_800X800.png",
	"group1_face_1080p.jpg",
	"rotate_4k.jpg",
	"face_special_1000X900.jpg",
	"kimage_sources/1920X1000.png",
	"kimage_sources/1920X1920.jpeg",
}

func TestComposeVideoNormal(t *testing.T) {
	client := NewAIClient(nil)
	client.ServiceInfo.Host = DEMO_HOST_STAGING
}
