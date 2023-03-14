package base

import (
	"io"
	"strings"
	"testing"
)

type DumbRequest struct {
	DumbId    string
	DumbInt   int32
	lowerCase string // should be ignored
}

type ImageSet struct {
	ImageURL    string `json:",omitempty"` // ImageURL is the url of the image
	ImageBase64 string `json:",omitempty"` // ImageBase64 is the base64 encoded image
	DataId      string
}

type DumbType struct {
	Name string
	Id   string
}
type DumbNestedRequest struct {
	DumbTypes []DumbType
	Dddd      string // Scenes is a comma separated string, e.g. "aSceneName,anotherSceneName,YetAnotherSceneName". For now only "porn" is supported
}

type DumbPanicType struct {
	Name  string
	AChan chan []string
}
type DumbPanicNestedRequest struct {
	DumbTypes DumbPanicType
}

type readerNestedStruct struct {
	aaa    string
	Reader io.Reader
}

type readerNestedStructNoField struct {
	aaa string
}

type readerNestedStructNoACorrectType struct {
	Reader string
}

func TestGetReaderFromStruct(t *testing.T) {
	a := readerNestedStruct{"", strings.NewReader("aaa")}
	reader, err := getReaderFromStruct(a)
	if err != nil {
		t.Fatalf("%e", err)
	} else {
		t.Logf("%T:%+v", reader, reader)
	}

	b := readerNestedStructNoField{""}
	reader, err = getReaderFromStruct(b)
	if err != nil {
		t.Logf("expected %e", err)
	} else {
		t.Fatalf("unexpected %T:%+v", reader, reader)
	}

	c := readerNestedStructNoACorrectType{""}
	reader, err = getReaderFromStruct(c)
	if err != nil {
		t.Logf("expected %e", err)
	} else {
		t.Fatalf("unexpected %T:%+v", reader, reader)
	}
}

func TestConvertStructToUrlValues(t *testing.T) {
	d := DumbRequest{"ddd", 123, "lowercase"}
	u, e := convertStructToUrlValues(d)
	if e != nil {
		t.Fatalf("%e", e)
	} else {
		t.Logf("%v", u.Encode())
	}
}

func TestConvertNestedStructToUrlValues(t *testing.T) {
	d := DumbNestedRequest{[]DumbType{{"apple", "1"}, {"banana", "234"}, {"orange", "345"}}, "ddd"}
	u, e := convertStructToUrlValues(d)
	if e != nil {
		t.Fatalf("%e", e)
	} else {
		t.Logf("%v", u.Encode())
	}
}

func TestConvertNestedStructPtrToUrlValues(t *testing.T) {
	d := DumbNestedRequest{[]DumbType{{"apple", "1"}, {"banana", "234"}, {"orange", "345"}}, "ddd"}
	u, e := convertStructToUrlValues(&d)
	if e != nil {
		t.Fatalf("%e", e)
	} else {
		t.Logf("%v", u.Encode())
	}
}

func TestConvertNestedWithPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	d := DumbPanicNestedRequest{DumbPanicType{"apple", nil}}
	u, e := convertStructToUrlValues(d)
	if e != nil {
		t.Fatalf("%e", e)
	} else {
		t.Logf("%v", u.Encode())
	}
}
