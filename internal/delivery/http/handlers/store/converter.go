package store

import "github.com/Nikkoz/mp.gateway/internal/domain/store"

func ToResponse(store *store.Store) *Response {
	return &Response{
		ID: store.ID,
		Short: Short{
			Name:        store.Name.String(),
			Marketplace: store.Marketplace.Uint8(),
		},
	}
}
