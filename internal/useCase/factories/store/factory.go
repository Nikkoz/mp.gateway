package store

import (
	"github.com/Nikkoz/mp.gateway/internal/useCase/adapters/grpc"
	"github.com/Nikkoz/mp.gateway/internal/useCase/adapters/storage"
)

type (
	Factory struct {
		adapterStorage storage.Store
		adapterGrpc    grpc.Store

		options Options
	}

	Options struct{}
)

func New(s storage.Store, g grpc.Store, o Options) *Factory {
	factory := &Factory{
		adapterStorage: s,
		adapterGrpc:    g,
	}

	factory.SetOption(o)

	return factory
}

func (f *Factory) SetOption(options Options) {
	if f.options != options {
		f.options = options
	}
}
