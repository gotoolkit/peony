package peony

type options struct {
	port       int
	configFile string
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithPort(port int) Option {
	return optionFunc(func(o *options) {
		o.port = port
	})
}

func WithConfigFile(path string) Option {
	return optionFunc(func(o *options) {
		o.configFile = path
	})
}
