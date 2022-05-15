package config

import (
	_ "embed"
	"log"
	"testing"
)

//go:embed test_data/config.yaml
var yamlContent []byte

type TestRedisConfig struct {
	Addr    string
	TimeOut string
}

func TestLocalConfig(t *testing.T) {
	log.Println("yamlContent: ", string(yamlContent))
	config, _, err := NewConfig(LocalConfig, WithYamlContent(yamlContent))

	testRedisconfig := new(TestRedisConfig)
	if err == nil {
		log.Println("config:", config)
		log.Println("config result:", config.Map())
		log.Println("Scan result:", config.Scan(testRedisconfig))
		log.Println("testRedisconfig result:", testRedisconfig)
	} else {
		log.Println("err: ", err)
	}

}
