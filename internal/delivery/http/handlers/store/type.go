package store

type (
	Response struct {
		ID uint `json:"id" binding:"required"`

		Short
	}

	Full struct {
		Short

		CampaignID   *uint64 `json:"campaign_id,omitempty"`
		ClientID     string  `json:"client_id" binding:"required"`
		ClientSecret string  `json:"client_secret" binding:"required"`
		Token        *string `json:"token,omitempty"`
		AuthToken    *string `json:"auth_token,omitempty"`
	}

	Short struct {
		Name        string `json:"name" binding:"required"`
		Marketplace uint8  `json:"marketplace" binding:"required"`
	}
)
