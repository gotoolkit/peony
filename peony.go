package peony

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
)

const defaultConfig = "."

type Core struct {
	context context.Context
	cancel  context.CancelFunc
	config  *viper.Viper
	server  *http.Server
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
	v.WatchConfig()
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(options.context)
	c := &Core{
		context: ctx,
		cancel:  cancel,
		config:  v,
		server: &http.Server{
			Addr: v.GetString("address"),
		},
	}

	return c, nil
}

func (c *Core) Start() error {
	c.run()
	log.Println(c.config.GetString("name"))
	go func() {
		for {
			select {
			case <-time.Tick(time.Second):
				log.Println(c.config.Get("name"))
			case <-c.context.Done():
				log.Println("done")
			}
		}
	}()
	// service connections
	if err := c.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
	return nil
}

func (c *Core) run() {
	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)

	go func() {
		select {

		case <-term:
			c.Stop()
			log.Println("term")
		}
	}()
}

func (c *Core) Stop() {
	c.cancel()
	if err := c.server.Shutdown(c.context); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
