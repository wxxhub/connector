package values

import (
	"fmt"
	"time"
)

// Values is returned by the reader
type Values interface {
	Bytes() ([]byte, error)
	Get(path ...string) Value
	Set(val interface{}, path ...string)
	Del(path ...string)
	Map() (map[string]interface{}, error)
	Scan(v interface{}) error
}

// Value represents a value of any type
type Value interface {
	Bool(def bool) (bool, error)
	Int(def int) (int, error)
	Float64(def float64) (float64, error)
	Bytes() ([]byte, error)
	String(def string) (string, error)
	StringSlice(def []string) ([]string, error)
	StringMap(def map[string]string) (map[string]string, error)
	Duration(def time.Duration) (time.Duration, error)
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
