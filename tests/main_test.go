package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabrielsilper/golang-gin/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r
}

func TestApiLiveController(t *testing.T) {
	// Arrange
	routerTester := SetupTest()
	routerTester.GET("/live", routes.LiveApi)
	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/live", nil)
	expectedBody := "Server is live..."

	// Act
	routerTester.ServeHTTP(resp, req)
	resultBody, _ := io.ReadAll(resp.Body)

	// Assert
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, expectedBody, string(resultBody))
}

func TestFindAllController(t *testing.T) {
	//Preciso aprender a mockar as funções do meu service
}
