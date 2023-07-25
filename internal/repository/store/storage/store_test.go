package storage

import (
	"github.com/Nikkoz/mp.gateway/internal/domain/store"
	"github.com/Nikkoz/mp.gateway/pkg/store/db"
	"github.com/Nikkoz/mp.gateway/pkg/types/context"
	"github.com/Nikkoz/mp.gateway/util"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"testing"
)

const StoreId uint = 1

var (
	connection *gorm.DB
	repository *Repository
	newStore   *store.Store
)

func TestMain(m *testing.M) {
	cancel := connectionDB()
	defer cancel()

	err := connection.AutoMigrate(&store.Store{})
	if err != nil {
		panic(err)
	}

	repository = New(connection, Options{})

	os.Exit(m.Run())
}

func TestCreateStore(t *testing.T) {
	assertion := assert.New(t)
	ctx := arrange()

	t.Run("Create store into DB", func(t *testing.T) {
		result, err := repository.CreateStore(ctx, newStore)

		assertion.NoError(err)
		assertion.NotEmpty(result)
		assertion.Equal(newStore.ID, result.ID)
		assertion.Equal(newStore.Name, result.Name)
		assertion.Equal(newStore.Marketplace, result.Marketplace)
	})
}

func arrange() context.Context {
	newStore = &store.Store{
		ID:          uint(util.RandomUInt(3, 5)),
		Name:        util.RandomStoreName(),
		Marketplace: "ozon",
	}

	return context.Empty()
}

func connectionDB() func() {
	conn, err := db.New(db.NewSettings(
		"localhost",
		5436,
		"marketplaces_test",
		"marketplaces",
		"secret",
		"disable",
		"mp_",
		4,
		1000,
	))
	if err != nil {
		panic(err)
	}

	connection = conn

	return func() {
		sqlDB, _ := conn.DB()
		err := sqlDB.Close()
		if err != nil {
			panic(err)
		}
	}
}
