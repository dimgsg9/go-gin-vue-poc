package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dimgsg9/booker_proto/backend/handler"
	"github.com/dimgsg9/booker_proto/backend/repository"
	"github.com/dimgsg9/booker_proto/backend/service"
	"github.com/gin-gonic/gin"
)

func inject(d *dataSources) (*gin.Engine, error) {
	log.Println("Injecting data sources")

	/*
	 * repository layer
	 */
	userRepository := repository.NewUserRepository(d.DB)
	tokenRepository := repository.NewTokenRepository(d.RedisClient)

	/*
	 * service layer
	 */
	userService := service.NewUserService(&service.USConfig{
		UserRepository: userRepository,
	})

	// load rsa keys
	privKeyFile := os.Getenv("PRIV_KEY_FILE")
	priv, err := ioutil.ReadFile(privKeyFile)

	if err != nil {
		return nil, fmt.Errorf("could not read private key pem file: %w", err)
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(priv)

	if err != nil {
		return nil, fmt.Errorf("could not parse private key: %w", err)
	}

	pubKeyFile := os.Getenv("PUB_KEY_FILE")
	pub, err := ioutil.ReadFile(pubKeyFile)

	if err != nil {
		return nil, fmt.Errorf("could not read public key pem file: %w", err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pub)

	if err != nil {
		return nil, fmt.Errorf("could not parse public key: %w", err)
	}

	// load refresh token secret from env variable
	refreshSecret := os.Getenv("REFRESH_SECRET")
	idTokenExp := os.Getenv("ID_TOKEN_EXP")
	refreshTokenExp := os.Getenv("REFRESH_TOKEN_EXP")

	idExp, err := strconv.ParseInt(idTokenExp, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse ID_TOKEN_EXP as int64: %w", err)
	}

	refreshExp, err := strconv.ParseInt(refreshTokenExp, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse REFRESH_TOKEN_EXP as int64: %w", err)
	}

	tokenService := service.NewTokenService(&service.TSConfig{
		TokenRepository:       tokenRepository,
		PrivKey:               privKey,
		PubKey:                pubKey,
		RefreshSecret:         refreshSecret,
		IDExpirationSecs:      idExp,
		RefreshExpirationSecs: refreshExp,
	})

	oauth2Creds := service.OAuthCreds{
		Cid:     os.Getenv("GOOGLE_OAUTH2_CLIENT_ID"),
		Csecret: os.Getenv("GOOGLE_OAUTH2_CLIENT_SECRET"),
	}

	oauthService := service.NewOAuthService(&service.OSConfig{
		Creds: oauth2Creds,
		Store: *d.SessionStore,
	})

	// initialize gin.Engine
	router := gin.Default()

	baseURL := os.Getenv("ACCOUNT_API_URL")

	// read in HANDLER_TIMEOUT
	handlerTimeout := os.Getenv("HANDLER_TIMEOUT")
	ht, err := strconv.ParseInt(handlerTimeout, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse HANDLER_TIMEOUT as int: %w", err)
	}

	handler.NewHandler(&handler.Config{
		R:               router,
		UserService:     userService,
		TokenService:    tokenService,
		OAuthService:    oauthService,
		BaseURL:         baseURL,
		TimeoutDuration: time.Duration(time.Duration(ht) * time.Second),
	})

	return router, nil
}
