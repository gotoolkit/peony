package peony

import "context"

type options struct {
	configPath string
	context    context.Context
}

// Option overrides behavior of Connect.
type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithConfig(path string) Option {
	return optionFunc(func(o *options) {
		o.configPath = path
	})
}

func WithContext(context context.Context) Option {
	return optionFunc(func(o *options) {
		o.context = context
	})
}
