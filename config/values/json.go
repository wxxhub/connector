package values

import (
	simple "github.com/bitly/go-simplejson"
	"time"
)

type JsonValues struct {
	content []byte
	sj      *simple.Json
}

type JsonValue struct {
	*simple.Json
}

func newJsonValues(content []byte) (Values, error) {
	sj := simple.New()
	if err := sj.UnmarshalJSON(content); err != nil {
		return nil, err
	}
	return &JsonValues{content: content, sj: sj}, nil
}

// YamlValues

func (j *JsonValues) Bytes() []byte {
	return j.content
}

func (j *JsonValues) Get(path ...string) Value {
	return nil
}

func (j *JsonValues) Set(val interface{}, path ...string) {

}

func (j *JsonValues) Del(path ...string) {

}

func (j *JsonValues) Map() map[string]interface{} {

	return nil
}

func (j *JsonValues) Scan(v interface{}) error {

	return nil
}

// YamlValue

func (j *JsonValue) Bool(def bool) bool {
	return def
}

func (j *JsonValue) Int(def int) int {
	return def
}

func (j *JsonValue) Float64(def float64) float64 {
	return def
}

func (j *JsonValue) String(def string) string {
	return def
}

func (j *JsonValue) Bytes() []byte {
	return nil
}

func (j *JsonValue) StringSlice(def []string) []string {
	//return
	return nil
}

func (j *JsonValue) StringMap(def map[string]string) map[string]string {
	return nil
}

func (j *JsonValue) Duration(def time.Duration) time.Duration {
	return def
}

func (j *JsonValue) Scan(v interface{}) error {

	return nil
}
