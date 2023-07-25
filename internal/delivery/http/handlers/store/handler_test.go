package store

import (
	"bytes"
	"encoding/json"
	"github.com/Nikkoz/mp.gateway/internal/domain/store"
	"github.com/Nikkoz/mp.gateway/internal/entities"
	mockUseCase "github.com/Nikkoz/mp.gateway/internal/useCase/mock"
	"github.com/Nikkoz/mp.gateway/pkg/types/context"
	enum "github.com/Nikkoz/mp.gateway/pkg/types/marketplace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"testing"
)

const StoreId uint = 1

var (
	uc      = new(mockUseCase.Store)
	handler *Handler
)

func TestCreate(t *testing.T) {
	tests := map[string]enum.Marketplace{
		//"Prepare Ozon data": enum.Ozon,
		"Prepare YM data": enum.YandexMarket,
	}

	handler = New(uc)

	for name, mp := range tests {
		assertion := arrangeHandler(t, mp)

		t.Run(name, func(t *testing.T) {
			handler.Create(ctx)

			var response Response
			_ = json.Unmarshal([]byte(recorder.Body.String()), &response)

			assertion.Equal(http.StatusCreated, recorder.Code)
			assertion.Equal(StoreId, response.ID)
			assertion.Equal(requestStore.Name, response.Name)
			assertion.Equal(requestStore.Marketplace, response.Marketplace)
		})
	}
}

func arrangeHandler(t *testing.T, mp enum.Marketplace) *assert.Assertions {
	assertion := assert.New(t)
	generateDataForJson(mp)

	uc.
		On(
			"Create",
			mock.Anything,
			mock.AnythingOfType("*store.Store"),
			mock.AnythingOfType("*entities.Access"),
		).
		Return(func(c context.Context, store *store.Store, access *entities.Access) *store.Store {
			assertion.Equal(requestStore.Name, store.Name.String())
			assertion.Equal(requestStore.Marketplace, store.Marketplace.Uint8())
			assertion.Equal(requestStore.ClientID, access.ClientID)
			assertion.Equal(requestStore.ClientSecret, access.ClientSecret)

			if mp.IsYandexMarket() {
				assertion.Equal(requestStore.Token, access.Token)
				assertion.Equal(requestStore.AuthToken, access.AuthToken)
			}

			store.ID = StoreId

			return store
		}, func(c context.Context, store *store.Store, access *entities.Access) error {
			return nil
		})

	marshal, err := json.Marshal(requestStore)
	if err != nil {
		panic(err)
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(marshal))

	return assertion
}
