package values

import "github.com/ghodss/yaml"

type yamlValues struct {
	jsonValues
}

//type YamlValue struct {
//	content []byte
//	paths   []string
//}

func newYamlValues(content []byte) (Values, error) {
	jsonContent, err := yaml.YAMLToJSON(content)
	if err != nil {
		return nil, err
	}
	return newJsonValues(jsonContent)
}

/*
// YamlValues

func (y *YamlValues) Bytes() []byte {
	return y.content
}

func (y *YamlValues) Get(path ...string) Value {
	return nil
}

func (y *YamlValues) Set(val interface{}, path ...string) {

}

func (y *YamlValues) Del(path ...string) {

}

func (y *YamlValues) Map() map[string]interface{} {

	return nil
}

func (y *YamlValues) Scan(v interface{}) error {

	return nil
}

// YamlValue

func (y *YamlValue) Bool(def bool) bool {
	return def
}

func (y *YamlValue) Int(def int) int {
	return def
}

func (y *YamlValue) Float64(def float64) float64 {
	return def
}

func (y *YamlValue) String(def string) string {
	return def
}

func (y *YamlValue) Bytes() []byte {
	return nil
}

func (y *YamlValue) StringSlice(def []string) []string {
	//return
	return nil
}

func (y *YamlValue) StringMap(def map[string]string) map[string]string {
	return nil
}

func (y *YamlValue) Duration(def time.Duration) time.Duration {
	return def
}

func (y *YamlValue) Scan(v interface{}) error {

	return nil
}

*/
