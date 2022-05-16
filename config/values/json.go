package values

import (
	"encoding/json"
	"fmt"
	simple "github.com/bitly/go-simplejson"
	"strconv"
	"strings"
	"time"
)

type jsonValues struct {
	content []byte
	sj      *simple.Json
}

type jsonValue struct {
	*simple.Json
}

func newJsonValues(content []byte) (Values, error) {
	sj := simple.New()
	if err := sj.UnmarshalJSON(content); err != nil {
		return nil, err
	}
	return &jsonValues{content: content, sj: sj}, nil
}

// YamlValues

func (j *jsonValues) Bytes() ([]byte, error) {
	return j.sj.MarshalJSON()
}

func (j *jsonValues) Get(path ...string) Value {
	return &jsonValue{j.sj.GetPath(path...)}
}

func (j *jsonValues) Set(val interface{}, path ...string) {
	j.sj.SetPath(path, val)
}

func (j *jsonValues) Del(path ...string) {
	// delete the tree?
	if len(path) == 0 {
		j.sj = simple.New()
		return
	}

	if len(path) == 1 {
		j.sj.Del(path[0])
		return
	}

	vals := j.sj.GetPath(path[:len(path)-1]...)
	vals.Del(path[len(path)-1])
	j.sj.SetPath(path[:len(path)-1], vals.Interface())
	return
}

func (j *jsonValues) Map() (map[string]interface{}, error) {
	return j.sj.Map()
}

func (j *jsonValues) Scan(v interface{}) error {
	b, err := j.sj.MarshalJSON()
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

// YamlValue

func (j *jsonValue) Bool(def bool) (bool, error) {
	b, err := j.Json.Bool()
	if err == nil {
		return def, err
	}

	str, ok := j.Interface().(string)
	if !ok {
		return def, err
	}

	b, err = strconv.ParseBool(str)
	if err != nil {
		return def, err
	}

	return b, nil
}

func (j *jsonValue) Int(def int) (int, error) {
	i, err := j.Json.Int()
	if err == nil {
		return i, nil
	}

	str, ok := j.Interface().(string)
	if !ok {
		return def, fmt.Errorf("failed")
	}

	i, err = strconv.Atoi(str)
	if err != nil {
		return def, err
	}

	return i, nil
}

func (j *jsonValue) Float64(def float64) (float64, error) {
	f, err := j.Json.Float64()
	if err == nil {
		return f, nil
	}

	str, ok := j.Interface().(string)
	if !ok {
		return def, fmt.Errorf("failed")
	}

	f, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return def, err
	}

	return f, nil
}

func (j *jsonValue) String(def string) (string, error) {
	return j.Json.MustString(def), nil
}

func (j *jsonValue) Bytes() ([]byte, error) {
	b, err := j.Json.Bytes()
	if err != nil {
		// try return marshalled
		b, err = j.Json.MarshalJSON()
		if err != nil {
			return []byte{}, err
		}
		return b, nil
	}
	return b, nil
}

func (j *jsonValue) StringSlice(def []string) ([]string, error) {
	v, err := j.Json.String()
	if err == nil {
		sl := strings.Split(v, ",")
		if len(sl) > 0 {
			return sl, nil
		}
	}
	return j.Json.MustStringArray(def), nil
}

func (j *jsonValue) StringMap(def map[string]string) (map[string]string, error) {
	m, err := j.Json.Map()
	if err != nil {
		return def, err
	}

	res := map[string]string{}

	for k, v := range m {
		res[k] = fmt.Sprintf("%v", v)
	}

	return res, nil
}

func (j *jsonValue) Duration(def time.Duration) (time.Duration, error) {
	v, err := j.Json.String()
	if err != nil {
		return def, err
	}

	value, err := time.ParseDuration(v)
	if err != nil {
		return def, err
	}

	return value, nil
}

func (j *jsonValue) Scan(v interface{}) error {
	b, err := j.Json.MarshalJSON()
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}
