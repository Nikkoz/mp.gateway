package marketplace

import "github.com/pkg/errors"

const (
	Ozon Marketplace = iota + 1
	YandexMarket
)

var ErrWrongValue = errors.Errorf("store's marketplace must be one of %d or %d", Ozon, YandexMarket)

type Marketplace uint8

func New(mp uint8) (*Marketplace, error) {
	newMp := Marketplace(mp)

	if newMp != Ozon && newMp != YandexMarket {
		return nil, ErrWrongValue
	}

	return &newMp, nil
}

func (m Marketplace) IsOzon() bool {
	return m == Ozon
}

func (m Marketplace) IsYandexMarket() bool {
	return m == YandexMarket
}

func (m Marketplace) String() string {
	if m.IsOzon() {
		return "ozon"
	}

	return "yandex_market"
}
