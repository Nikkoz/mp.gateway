package store

import (
	"bytes"
	"encoding/json"
	enum "github.com/Nikkoz/mp.gateway/pkg/types/marketplace"
	"github.com/Nikkoz/mp.gateway/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

var (
	recorder     *httptest.ResponseRecorder
	ctx          *gin.Context
	request      *http.Request
	requestStore *Full
)

func TestMain(m *testing.M) {
	resetData()

	os.Exit(m.Run())
}

func resetData() {
	recorder = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(recorder)

	request = &http.Request{
		URL: &url.URL{},
	}

	ctx.Request = request
}

func TestMakeData(t *testing.T) {
	tests := map[string]enum.Marketplace{
		"Prepare Ozon data": enum.Ozon,
		"Prepare YM data":   enum.YandexMarket,
	}

	for name, mp := range tests {
		arrangePreparer(mp)

		t.Run(name, func(t *testing.T) {
			assertion := assert.New(t)

			store, access, err := makeData(ctx)
			if err != nil {
				panic(err)
			}

			assertion.Equal(requestStore.Name, store.Name.String())
			assertion.Equal(requestStore.Marketplace, store.Marketplace.Uint8())
			assertion.Equal(requestStore.ClientID, access.ClientID)
			assertion.Equal(requestStore.ClientSecret, access.ClientSecret)

			if mp.IsYandexMarket() {
				assertion.Equal(requestStore.Token, access.Token)
				assertion.Equal(requestStore.AuthToken, access.AuthToken)
			}
		})
	}
}

func arrangePreparer(mp enum.Marketplace) {
	generateDataForJson(mp)

	marshal, err := json.Marshal(requestStore)
	if err != nil {
		panic(err)
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(marshal))
}

func generateDataForJson(mp enum.Marketplace) {
	name := util.RandomStoreName()
	requestStore = &Full{
		Short: Short{
			Name:        name.String(),
			Marketplace: uint8(mp),
		},
	}

	if mp.IsYandexMarket() {
		campaignId := rand.Uint64()
		token := util.RandomStoreToken()
		authToken := util.RandomStoreAuthToken()

		requestStore.CampaignID = &campaignId
		requestStore.Token = &token
		requestStore.AuthToken = &authToken
	}

	requestStore.ClientID = util.RandomStoreClientId()
	requestStore.ClientSecret = util.RandomStoreClientSecret()
}
