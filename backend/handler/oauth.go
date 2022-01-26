package handler

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"

	"github.com/dimgsg9/booker_proto/backend/model/apperrors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetLoginURL(c *gin.Context) {

	session := sessions.Default(c)

	if c.Query("req") == "url" {
		ctx := c.Request.Context()
		state := randToken()

		session.Set("state", state)
		session.Save()

		url, _ := h.OAuthService.GetLoginURL(ctx, "google", state) //TODO error handling

		c.JSON(http.StatusOK, gin.H{
			"loginUrl": url,
		})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "uknown request",
		})
	}

}

func (h *Handler) AuthCallback(c *gin.Context) {
	session := sessions.Default(c)

	rstate := session.Get("state")
	if rstate != c.Query("state") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid session state",
		})
		return
	}

	code := c.Query("code")

	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid session code",
		})
		return
	}

	ctx := c.Request.Context()

	user, err := h.OAuthService.AuthCallback(ctx, "google", code) //should return User that will be then passsed to user service OauthLogin method
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "failed to retrieve the user",
		})
		return
	}

	err = h.UserService.OauthLogin(ctx, user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
		return
	}

	tokens, err := h.TokenService.NewPairFromUser(ctx, user, "") //TODO: what about previous token?
	if err != nil {
		log.Printf("Failed to create tokens for user: %v\n", err.Error())

		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tokens": tokens,
	})
	// c.JSON(http.StatusOK, gin.H{
	// 	"user": user, //TODO: temporary return value, should auth the user instead
	// })

}

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
