package store

import (
	"fmt"
	"github.com/Nikkoz/mp.gateway/internal/domain/store"
	"github.com/Nikkoz/mp.gateway/internal/domain/store/types/marketplace"
	"github.com/Nikkoz/mp.gateway/internal/domain/store/types/name"
	"github.com/Nikkoz/mp.gateway/internal/entities"
	"github.com/Nikkoz/mp.gateway/pkg/types/context"
	"github.com/Nikkoz/mp.gateway/pkg/types/logger"
	enum "github.com/Nikkoz/mp.gateway/pkg/types/marketplace"
	"github.com/gin-gonic/gin"
)

func makeData(c *gin.Context) (*store.Store, *entities.Access, error) {
	data := &Full{}
	if err := c.ShouldBindJSON(&data); err != nil {
		return nil, nil, fmt.Errorf("payload is not correct, Error: %w", err)
	}

	ctx := context.New(c)

	enumValue, err := enum.New(data.Marketplace)
	if err != nil {
		return nil, nil, logger.ErrorWithContext(ctx, err)
	}

	mp, err := marketplace.New(*enumValue)
	if err != nil {
		return nil, nil, logger.ErrorWithContext(ctx, err)
	}

	sName, err := name.New(data.Name)
	if err != nil {
		return nil, nil, logger.ErrorWithContext(ctx, err)
	}

	newStore := store.New(nil, *mp, *sName)
	access := &entities.Access{
		CampaignID:   data.CampaignID,
		ClientID:     data.ClientID,
		ClientSecret: data.ClientSecret,
		Token:        data.Token,
		AuthToken:    data.AuthToken,
	}

	return newStore, access, nil
}
