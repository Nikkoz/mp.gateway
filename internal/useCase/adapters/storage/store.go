package storage

import (
	"github.com/Nikkoz/mp.gateway/internal/domain/store"
	"github.com/Nikkoz/mp.gateway/pkg/types/context"
)

type (
	Store interface {
		CreateStore(ctx context.Context, store *store.Store) (*store.Store, error)
	}
)
