package handler

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetLoginURL(c *gin.Context) {
	if c.Query("req") == "url" {
		ctx := c.Request.Context()
		state := randToken()
		url, _ := h.OAuthService.GetLoginURL(ctx, "google", state)

		c.JSON(http.StatusOK, gin.H{
			"loginUrl": url,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
	}

}

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
