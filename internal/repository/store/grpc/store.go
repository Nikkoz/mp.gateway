package grpc

import (
	"github.com/Nikkoz/mp.gateway/internal/domain/store"
	"github.com/Nikkoz/mp.gateway/internal/entities"
	"github.com/Nikkoz/mp.gateway/pkg/types/context"
	"github.com/Nikkoz/mp.gateway/pkg/types/logger"
)

func (r *Repository) Create(ctx context.Context, store *store.Store, access *entities.Access) (uint, error) {
	response, err := r.client.CreateStore(ctx, ToRequest(store, access))

	if err != nil {
		return 0, logger.ErrorWithContext(ctx, err)
	}

	return uint(response.GetId().GetValue()), nil
}
