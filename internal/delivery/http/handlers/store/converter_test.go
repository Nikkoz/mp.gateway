package store

import (
	domain "github.com/Nikkoz/mp.gateway/internal/domain/store"
	"github.com/Nikkoz/mp.gateway/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToResponse(t *testing.T) {
	assertion := assert.New(t)
	store := arrangeConverter()

	t.Run("Convert response", func(t *testing.T) {
		response := ToResponse(store)

		assertion.Equal(store.ID, response.ID)
		assertion.Equal(store.Marketplace.Uint8(), response.Short.Marketplace)
		assertion.Equal(store.Name.String(), response.Short.Name)
	})
}

func arrangeConverter() *domain.Store {
	return &domain.Store{
		ID:          1,
		Marketplace: "ozon",
		Name:        util.RandomStoreName(),
	}
}
