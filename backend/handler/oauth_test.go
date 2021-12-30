package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dimgsg9/booker_proto/backend/model/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestOAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockOAuthService := new(mocks.MockOAuthService)

	router := gin.Default()

	NewHandler(&Config{
		R:            router,
		OAuthService: mockOAuthService,
	})

	t.Run("Bad request", func(t *testing.T) {
		rr := httptest.NewRecorder()

		url := "https://accounts.google.com/o/oauth2/v2/auth?scope=https%3A//www.googleapis.com/auth/drive.metadata.readonly&access_type=offline&include_granted_scopes=true&response_type=code&state=state_parameter_passthrough_value&redirect_uri=https%3A//oauth2.example.com/code&client_id=client_id"

		mockOAuthService.On("GetLoginURL", mock.AnythingOfType("*context.emptyCtx"), "google", mock.AnythingOfType("string")).Return(url, nil)

		request, err := http.NewRequest(http.MethodGet, "/oauth?foo=bar", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		mockOAuthService.AssertNotCalled(t, "GetLoginURL")

	})

	t.Run("Success", func(t *testing.T) {
		rr := httptest.NewRecorder()

		url := "https://accounts.google.com/o/oauth2/v2/auth?scope=https%3A//www.googleapis.com/auth/drive.metadata.readonly&access_type=offline&include_granted_scopes=true&response_type=code&state=state_parameter_passthrough_value&redirect_uri=https%3A//oauth2.example.com/code&client_id=client_id"

		mockOAuthService.On("GetLoginURL", mock.AnythingOfType("*context.emptyCtx"), "google", mock.AnythingOfType("string")).Return(url, nil)

		request, err := http.NewRequest(http.MethodGet, "/oauth?req=url", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"loginUrl": url,
		})

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockOAuthService.AssertExpectations(t)

	})

}
