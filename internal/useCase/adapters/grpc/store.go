package grpc

import (
	"github.com/Nikkoz/mp.gateway/internal/domain/store"
	"github.com/Nikkoz/mp.gateway/internal/entities"
	"github.com/Nikkoz/mp.gateway/pkg/types/context"
)

type (
	Store interface {
		Create(ctx context.Context, store *store.Store, access *entities.Access) (uint, error)
	}
)
