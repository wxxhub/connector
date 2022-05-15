package config

type options struct {
	// yaml
	yamlContent []byte
}

type Option func(*options)

func WithYamlContent(content []byte) Option {
	return func(o *options) {
		o.yamlContent = content
	}
}
