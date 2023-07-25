package store

import (
	"github.com/Nikkoz/mp.gateway/internal/domain/store/types/marketplace"
	"github.com/Nikkoz/mp.gateway/internal/domain/store/types/name"
)

type (
	Store struct {
		ID uint `gorm:"primaryKey;AUTO_INCREMENT"`

		Marketplace marketplace.Marketplace `gorm:"not null;comment:Тип маркетплейса" sql:"type:ENUM(ozon, yandex_market)"`
		Name        name.Name               `gorm:"size:100;not null;comment:Название магазина"`
	}
)

func New(id *uint, mp marketplace.Marketplace, name name.Name) *Store {
	store := &Store{
		Marketplace: mp,
		Name:        name,
	}

	if id != nil {
		store.ID = *id
	}

	return store
}
