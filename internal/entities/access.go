package entities

type Access struct {
	CampaignID   *uint64
	ClientID     string
	ClientSecret string
	Token        *string
	AuthToken    *string
}
