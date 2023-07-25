package storage

import (
	"github.com/Nikkoz/mp.gateway/internal/domain/store"
	"github.com/Nikkoz/mp.gateway/pkg/types/context"
	"github.com/Nikkoz/mp.gateway/pkg/types/logger"
)

func (r *Repository) CreateStore(c context.Context, store *store.Store) (*store.Store, error) {
	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	if err := r.db.Create(&store).Error; err != nil {
		return nil, logger.ErrorWithContext(ctx, err)
	}

	return store, nil
}
