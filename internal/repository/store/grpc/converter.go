package grpc

import (
	"github.com/Nikkoz/mp.gateway/internal/domain/store"
	"github.com/Nikkoz/mp.gateway/internal/domain/store/types/marketplace"
	"github.com/Nikkoz/mp.gateway/internal/entities"
	"github.com/Nikkoz/mp.product_service/pkg/protobuf/stores"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ToRequest(store *store.Store, access *entities.Access) *stores.CreateStoreRequest {
	request := &stores.Store{
		Marketplace:  getMarketplace(store.Marketplace),
		Name:         store.Name.String(),
		ClientId:     access.ClientID,
		ClientSecret: access.ClientSecret,
	}

	if access.CampaignID != nil {
		request.CampaignId = &wrapperspb.Int64Value{
			Value: int64(*access.CampaignID),
		}
	}

	if access.Token != nil {
		request.Token = &wrapperspb.StringValue{
			Value: *access.Token,
		}
	}

	if access.AuthToken != nil {
		request.AuthToken = &wrapperspb.StringValue{
			Value: *access.AuthToken,
		}
	}

	return &stores.CreateStoreRequest{
		Store: request,
	}
}

func getMarketplace(mp marketplace.Marketplace) stores.Store_Marketplace {
	if mp == "ozon" {
		return stores.Store_Ozon
	}

	return stores.Store_YandexMarket
}
