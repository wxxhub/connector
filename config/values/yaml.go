package values

import "gopkg.in/yaml.v2"

type Yaml struct {
	content []byte
	paths   []string
}

func newYamlValues(content []byte) *Yaml {
	return &Yaml{content: content}
}

func (y *Yaml) Bool(def bool) bool {
	return def
}

func (y *Yaml) Int(def int) int {
	return def
}

func (y *Yaml) Float64(def float64) float64 {
	return def
}

func (y *Yaml) Bytes(def []byte) []byte {
	return def
}

func (y *Yaml) String(def string) string {
	return def
}

func (y *Yaml) StringSlice(def []string) []string {
	return def
}

func (y *Yaml) Map() map[string]interface{} {
	m := make(map[string]interface{})
	yaml.Unmarshal(y.content, m)
	return m
}

func (y *Yaml) StringMap() map[string]string {
	return nil
}

func (y *Yaml) Get(path ...string) Values {
	return y
}

func (y *Yaml) Scan(v interface{}) error {
	return nil
}
