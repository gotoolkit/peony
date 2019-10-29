package peony

import (
	"context"

	"github.com/spf13/viper"
)

const defaultConfig = "."

type Core struct {
	context context.Context
	config  *viper.Viper
}

func New(opts ...Option) (*Core, error) {
	options := options{
		context:    context.Background(),
		configPath: defaultConfig,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	v := viper.New()
	v.AddConfigPath(options.configPath)
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := &Core{
		context: options.context,
		config:  v,
	}

	return c, nil
}
