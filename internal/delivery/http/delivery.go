package http

import (
	"fmt"
	"github.com/Nikkoz/mp.gateway/internal/configs"
	"github.com/Nikkoz/mp.gateway/internal/delivery/http/handlers/store"
	"github.com/Nikkoz/mp.gateway/internal/useCase/interfaces"
	"github.com/gin-gonic/gin"
)

type (
	Delivery struct {
		ucStore interfaces.Store

		router *gin.Engine

		options  Options
		handlers Handlers
	}

	Options struct {
		Notify chan error
	}

	Handlers struct {
		Store *store.Handler
	}
)

func New(ucStore interfaces.Store, o Options) *Delivery {
	d := &Delivery{
		ucStore: ucStore,
	}

	d.setOptions(o)
	d.setHandlers()

	return d
}

func (d *Delivery) setOptions(options Options) {
	if options.Notify == nil {
		d.options.Notify = make(chan error, 1)
	}

	if d.options != options {
		d.options = options
	}
}

func (d *Delivery) setHandlers() {
	d.handlers = Handlers{
		Store: store.New(d.ucStore),
	}
}

func (d *Delivery) Run(config configs.Config) {
	d.initRouter(config)

	go func() {
		defer close(d.options.Notify)

		d.options.Notify <- d.router.Run(fmt.Sprintf("%s:%d", config.Http.Host, config.Http.Port))
	}()
}

func (d *Delivery) Notify() <-chan error {
	return d.options.Notify
}
