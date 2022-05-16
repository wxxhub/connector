package config

import (
	_ "embed"
	"log"
	"testing"
)

//go:embed test_data/config.yaml
var yamlContent []byte

type TestRedisConfig struct {
	Addr    string `json:"addr"`
	TimeOut int    `json:"time_out"`
}

func TestLocalConfig(t *testing.T) {
	log.Println("yamlContent: ", string(yamlContent))
	config, _, err := NewConfig(LocalConfig, WithYamlContent(yamlContent))

	testRedisconfig := new(TestRedisConfig)
	if err == nil {
		log.Println("config:", config)
		m, err := config.Map()
		log.Println("config result m:", m)
		log.Println("config result err:", err)
		log.Println("Scan result:", config.Scan(testRedisconfig))
		log.Println("Scan result:", config.Get("redis").Scan(testRedisconfig))
		log.Println("testRedisconfig result:", testRedisconfig)
	} else {
		log.Println("err: ", err)
	}

}
