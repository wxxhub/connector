package config

import (
	"fmt"
	"github.com/wxxhub/middleware/config/values"
	"sync"
)

type localConfig struct {
	sync.RWMutex

	yamlContent []byte
	vales       values.Values
}

func newLocalConfig(opts ...Option) (Config, Watcher, error) {
	opt := &options{}
	for _, o := range opts {
		o(opt)
	}

	if len(opt.yamlContent) > 0 {

		vales, err := values.NewValues(values.Yml, opt.yamlContent)
		if err != nil {
			return nil, nil, err
		}

		c := &localConfig{
			yamlContent: opt.yamlContent,
			vales:       vales,
		}

		return c, nil, nil
	} else {
		return nil, nil, fmt.Errorf("yamlContent is empty")
	}

	return nil, nil, fmt.Errorf("new local config failed")
}

func (l *localConfig) Map() (map[string]interface{}, error) {
	l.RLock()
	defer l.RUnlock()
	return l.vales.Map()
}

func (l *localConfig) Scan(v interface{}) error {
	l.RLock()
	defer l.RUnlock()
	return l.vales.Scan(v)
}

func (l *localConfig) Get(path ...string) values.Value {
	l.RLock()
	defer l.RUnlock()
	return l.vales.Get(path...)
}
