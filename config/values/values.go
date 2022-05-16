package values

import (
	"fmt"
	"time"
)

// Values is returned by the reader
type Values interface {
	Bytes() []byte
	Get(path ...string) Value
	Set(val interface{}, path ...string)
	Del(path ...string)
	Map() map[string]interface{}
	Scan(v interface{}) error
}

// Value represents a value of any type
type Value interface {
	Bool(def bool) bool
	Int(def int) int
	Float64(def float64) float64
	Bytes() []byte
	String(def string) string
	StringSlice(def []string) []string
	StringMap(def map[string]string) map[string]string
	Duration(def time.Duration) time.Duration
	Scan(val interface{}) error
}

type Type string

var (
	Yml  Type = "yml"
	Json Type = "json"
)

func NewValues(valueType Type, content []byte) (Values, error) {
	switch valueType {
	case Yml:
		return newYamlValues(content)
	case Json:
		return newJsonValues(content)
	default:
		return nil, fmt.Errorf("uncomplete %s", valueType)
	}
}
