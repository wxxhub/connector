package values

import "fmt"

type Values interface {
	Bool(def bool) bool
	Int(def int) int
	Float64(def float64) float64
	Bytes(def []byte) []byte
	String(def string) string
	StringSlice(def []string) []string
	Map() map[string]interface{}
	StringMap() map[string]string

	Get(path ...string) Values
	Scan(v interface{}) error
}

type Type string

var (
	Yml Type
)

func NewValues(valueType Type, content []byte) (Values, error) {
	switch valueType {
	case Yml:
		return newYamlValues(content), nil
	default:
		return nil, fmt.Errorf("uncomplete %s", valueType)
	}
}
