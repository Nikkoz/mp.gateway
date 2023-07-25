package grpc

import (
	"github.com/Nikkoz/mp.product_service/pkg/protobuf/stores"
	"google.golang.org/grpc"
)

type (
	Repository struct {
		client stores.StoreServiceClient

		options Options
	}

	Options struct{}
)

func New(conn *grpc.ClientConn, o Options) *Repository {
	repo := &Repository{
		client: stores.NewStoreServiceClient(conn),
	}

	repo.SetOptions(o)

	return repo
}

func (r *Repository) SetOptions(options Options) {
	if r.options != options {
		r.options = options
	}
}
