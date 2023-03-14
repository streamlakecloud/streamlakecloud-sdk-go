package base

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// the following three type and func are from src/encoding/json/tags.go

// tagOptions is the string following a comma in a struct field's "json"
// tag, or the empty string. It does not include the leading comma.
type tagOptions string

// parseTag splits a struct field's json tag into its name and
// comma-separated options.
func parseTag(tag string) (string, tagOptions) {
	tag, opt, _ := strings.Cut(tag, ",")
	return tag, tagOptions(opt)
}

// Contains reports whether a comma-separated list of options
// contains a particular substr flag. substr must be surrounded by a
// string boundary or commas.
func (o tagOptions) Contains(optionName string) bool {
	if len(o) == 0 {
		return false
	}
	s := string(o)
	for s != "" {
		var name string
		name, s, _ = strings.Cut(s, ",")
		if name == optionName {
			return true
		}
	}
	return false
}

func getReaderFromStruct(a interface{}) (io.Reader, error) {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Pointer {
		// get pointer underlying value
		v = reflect.ValueOf(a).Elem()
	}
	f := v.FieldByName("Reader")
	if !f.IsValid() || f.IsZero() {
		return nil, fmt.Errorf("not a valid field: Reader")
	}
	fv := f.Interface()
	readerType := reflect.TypeOf((*io.Reader)(nil)).Elem()
	if f.CanConvert(readerType) {
		return fv.(io.Reader), nil
	}
	return nil, fmt.Errorf("field Reader cannot convert to io.Reader type")
}

func convertStructToUrlValues(a interface{}) (url.Values, error) {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Pointer {
		// get pointer underlying value
		v = reflect.ValueOf(a).Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf(v.Kind().String() + " not a struct")
	}
	params := url.Values{}

	for i := 0; i < v.NumField(); i++ {
		fn := v.Type().Field(i).Name

		if v.Type().Field(i).PkgPath != "" {
			// unexported
			continue
		}

		f := v.Field(i)
		if !f.IsValid() {
			// invalid, zero values
			continue
		}

		// reuse json tag on field
		s := v.Type().Field(i).Tag.Get("json")
		name, options := parseTag(s)
		if name == "-" {
			// skip field
			continue
		} else if len(name) > 0 {
			fn = name
		}

		if options.Contains("omitempty") && f.IsZero() {
			// omit empty value
			continue
		}
		var o string
		if f.CanInt() {
			o = strconv.FormatInt(f.Int(), 10)
		} else if f.CanUint() {
			o = strconv.FormatUint(f.Uint(), 10)
		} else if f.CanFloat() {
			o = strconv.FormatFloat(f.Float(), 'f', -1, 64)
		} else if f.CanConvert(reflect.TypeOf("")) {
			fv := f.Convert(reflect.TypeOf("")).String()
			o = fv
		} else {
			fv := f.Interface()
			switch f.Kind() {
			case reflect.String:
				o = fv.(string)
			case reflect.Array, reflect.Slice, reflect.Map, reflect.Struct:
				bytes, err := json.Marshal(fv)
				if err != nil {
					panic(fmt.Sprintf("failed to marshal %v: %v", f, err))
				} else {
					o = string(bytes)
				}
			default:
				panic(fmt.Sprintf("unsupported type: %s", f.Kind()))
			}
		}
		params.Set(fn, o)
	}
	return params, nil
}

// headers with same key will be override by header2
func mergeHeader(header1, header2 http.Header) (header http.Header) {
	header = header1.Clone()
	if header == nil {
		header = http.Header{}
	}
	for k, v := range header2 {
		header.Set(k, strings.Join(v, ";"))
	}
	return
}
