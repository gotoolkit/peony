package peony

import "github.com/gin-gonic/gin"

type App struct {
	Router *gin.Engine
}

func New(opts ...Option) *App {

	options := options{}

	for _, o := range opts {
		o.apply(&options)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	a := &App{Router: r}
	return a
}

func (a *App) Run() error {
	return a.Router.Run()
}
