package service

import (
	"context"

	"github.com/dimgsg9/booker_proto/backend/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuthCreds struct {
	Cid     string
	Csecret string
}

type oauthService struct {
	Creds OAuthCreds
}

type OSConfig struct {
	Creds OAuthCreds
}

func NewOAuthService(c *OSConfig) model.OAuthService {
	return &oauthService{
		Creds: c.Creds,
	}
}

func (o *oauthService) GetLoginURL(ctx context.Context, provider string, state string) (string, error) {

	conf := &oauth2.Config{
		ClientID:     o.Creds.Cid,
		ClientSecret: o.Creds.Csecret,
		RedirectURL:  "http://localhost:9090/auth",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	url := conf.AuthCodeURL(state)

	return url, nil

}
