package store

import (
	"github.com/Nikkoz/mp.gateway/internal/domain/store"
	"github.com/Nikkoz/mp.gateway/internal/entities"
	"github.com/Nikkoz/mp.gateway/pkg/types/context"
)

func (f *Factory) Create(ctx context.Context, store *store.Store, access *entities.Access) (*store.Store, error) {
	id, err := f.adapterGrpc.Create(ctx, store, access)
	if err != nil {
		return nil, err
	}

	store.ID = id

	return f.adapterStorage.CreateStore(ctx, store)
}
