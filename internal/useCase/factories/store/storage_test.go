package store

import (
	"github.com/Nikkoz/mp.gateway/internal/domain/store"
	"github.com/Nikkoz/mp.gateway/internal/entities"
	mockStoreGrpc "github.com/Nikkoz/mp.gateway/internal/repository/store/grpc/mock"
	mockStoreStorage "github.com/Nikkoz/mp.gateway/internal/repository/store/storage/mock"
	"github.com/Nikkoz/mp.gateway/pkg/types/context"
	"github.com/Nikkoz/mp.gateway/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

const StoreId uint = 1

var (
	factory           *Factory
	storageRepository = new(mockStoreStorage.Store)
	grpcRepository    = new(mockStoreGrpc.Store)
	storeData         *store.Store
	accessData        *entities.Access
)

func TestCreate(t *testing.T) {
	assertion := arrangeCreate(t)

	t.Run("Create store", func(t *testing.T) {
		ctx := context.Empty()

		result, err := factory.Create(ctx, storeData, accessData)

		assertion.NoError(err)
		assertion.Equal(StoreId, result.ID)
		assertion.Equal(storeData.Name, result.Name)
		assertion.Equal(storeData.Marketplace, result.Marketplace)
	})
}

func arrangeCreate(t *testing.T) *assert.Assertions {
	generateStoreData()

	factory = New(storageRepository, grpcRepository, Options{})
	assertion := assert.New(t)

	grpcRepository.
		On(
			"Create",
			mock.Anything,
			mock.AnythingOfType("*store.Store"),
			mock.AnythingOfType("*entities.Access"),
		).
		Return(func(ctx context.Context, store *store.Store, access *entities.Access) uint {
			assertion.Equal(storeData.Name, store.Name)
			assertion.Equal(storeData.Marketplace, store.Marketplace)
			assertion.Equal(accessData.ClientID, access.ClientID)
			assertion.Equal(accessData.ClientSecret, access.ClientSecret)
			assertion.Equal(accessData.Token, access.Token)
			assertion.Equal(accessData.AuthToken, access.AuthToken)

			return StoreId
		}, func(ctx context.Context, store *store.Store, access *entities.Access) error {
			return nil
		})

	storageRepository.
		On(
			"CreateStore",
			mock.Anything,
			mock.AnythingOfType("*store.Store"),
		).
		Return(func(ctx context.Context, store *store.Store) *store.Store {
			assertion.Equal(storeData.Name, store.Name)
			assertion.Equal(storeData.Marketplace, store.Marketplace)

			store.ID = StoreId

			return store
		}, func(ctx context.Context, store *store.Store) error {
			return nil
		})

	return assertion
}

func generateStoreData() {
	storeData = &store.Store{
		Name:        util.RandomStoreName(),
		Marketplace: "yandex_market",
	}

	campaignId := util.RandomUInt(5, 7)
	token, authToken := util.RandomStoreToken(), util.RandomStoreAuthToken()

	accessData = &entities.Access{
		ClientID:     util.RandomStoreClientId(),
		ClientSecret: util.RandomStoreClientSecret(),
		CampaignID:   &campaignId,
		Token:        &token,
		AuthToken:    &authToken,
	}
}
