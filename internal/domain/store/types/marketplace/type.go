package marketplace

import enum "github.com/Nikkoz/mp.gateway/pkg/types/marketplace"

type Marketplace string

func New(mp enum.Marketplace) (*Marketplace, error) {
	m := Marketplace(mp.String())

	return &m, nil
}

func (mp *Marketplace) String() string {
	return string(*mp)
}

func (mp *Marketplace) Uint8() uint8 {
	if mp.String() == "ozon" {
		return uint8(enum.Ozon)
	}

	return uint8(enum.YandexMarket)
}
