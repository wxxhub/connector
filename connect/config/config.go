package config

import (
	"context"
	"fmt"
	"github.com/wxxhub/middleware/config"
	"sync"
)

var configs *configMap

type configMap struct {
	sync.RWMutex
	Map     map[string]config.Config
	Watcher map[string]config.Watcher
}

func init() {
	configs = new(configMap)
	configs.Map = make(map[string]config.Config)
}

func ConnectConfig(ctx context.Context, confType config.Type, confName string, opts ...config.Option) (config.Config, config.Watcher, error) {
	key := fmt.Sprintf("%s/%s", confType, confName)

	configs.RLock()
	_, ok := configs.Map[key]
	configs.RUnlock()

	if !ok {
		configs.Lock()
		defer configs.Unlock()

		_, ok := configs.Map[key]
		if !ok {
			conf, watcher, err := config.NewConfig(config.LocalConfig, opts...)
			if err == nil {
				configs.Map[key] = conf
				configs.Watcher[key] = watcher
			}
			return conf, watcher, err
		}
	}

	return configs.Map[key], configs.Watcher[key], nil
}
