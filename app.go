package peony

import (
	"fmt"

	"github.com/gotoolkit/config"

	"github.com/gin-gonic/gin"
	"github.com/gotoolkit/log"
)

type App struct {
	opt    *options
	Router *gin.Engine
}

func New(opts ...Option) *App {

	opt := options{
		port:       8080,
		configFile: "./config/peny.yml",
	}

	for _, o := range opts {
		o.apply(&opt)
	}

	config.Setup(
		config.WithAutoEnv(true),
		config.WithWatchEnable(true),
		config.WithFile(opt.configFile),
		config.WithDefault(map[string]interface{}{
			"server.port": opt.port,
		}))

	log.Setup()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	a := &App{Router: r, opt: &opt}
	return a
}

func (a *App) Run() error {
	return a.Router.Run(fmt.Sprintf(":%d", config.GetInt("server.port")))
}
