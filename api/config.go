package api

import (
	"os"

	"github.com/fosshostorg/teardrop/internal/pkg/models"
	"github.com/palantir/go-githubapp/githubapp"
)

func ReadEnv() (*models.Config, error) {
	c := models.Config{
		Github: githubapp.Config{
			WebURL:   "https://github.com",
			V3APIURL: "https://api.github.com/",
			App: struct {
				IntegrationID int64  "yaml:\"integration_id\" json:\"integrationId\""
				WebhookSecret string "yaml:\"webhook_secret\" json:\"webhookSecret\""
				PrivateKey    string "yaml:\"private_key\" json:\"privateKey\""
			}{
				IntegrationID: 163011,
				PrivateKey:    os.Getenv("GITHUB_PRIVATE_KEY"),
			},
			OAuth: struct {
				ClientID     string "yaml:\"client_id\" json:\"clientId\""
				ClientSecret string "yaml:\"client_secret\" json:\"clientSecret\""
			}{ClientID: "Iv1.e5e11f90d5f34a85",
				ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET")},
		},
	}

	return &c, nil
}
