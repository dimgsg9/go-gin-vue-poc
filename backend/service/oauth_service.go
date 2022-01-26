package service

import (
	"context"
	"encoding/json"
	"io"

	"github.com/dimgsg9/booker_proto/backend/model"
	"github.com/gin-contrib/sessions/redis"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuthCreds struct {
	Cid     string
	Csecret string
}

type oauthService struct {
	Creds OAuthCreds
	Store redis.Store
}

type OSConfig struct {
	Creds OAuthCreds
	Store redis.Store
}

func NewOAuthService(c *OSConfig) model.OAuthService {
	return &oauthService{
		Creds: c.Creds,
		Store: c.Store,
	}
}

//TODO: rename to get OauthURL()
func (o *oauthService) GetLoginURL(ctx context.Context, provider string, state string) (string, error) {

	conf := &oauth2.Config{
		ClientID:     o.Creds.Cid,
		ClientSecret: o.Creds.Csecret,
		RedirectURL:  "http://booker.app/api/account/oauth/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	url := conf.AuthCodeURL(state)

	return url, nil //TODO: error return?

}

func (o *oauthService) AuthCallback(ctx context.Context, provider string, code string) (*model.User, error) {

	user := &model.User{}
	var error error

	conf := &oauth2.Config{
		ClientID:     o.Creds.Cid,
		ClientSecret: o.Creds.Csecret,
		RedirectURL:  "http://booker.app/api/account/oauth/callback", //TODO: make it configurable from .env
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		error = err
	}

	client := conf.Client(ctx, tok)

	res, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		error = err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		error = err
	}

	json.Unmarshal(data, &user)

	return user, error

}

func (o *oauthService) GetSessionStore() (redis.Store, error) {
	return o.Store, nil //TODO: proper error return
}
