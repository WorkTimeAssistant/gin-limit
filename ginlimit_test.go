package ginlimit

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestUnderLimit(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(Limit(100, 100))
	router.GET("/test", func(c *gin.Context) {
		c.String(200, "OK")
	})

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, r)

	//assertions here
	assert.Equal(200, w.Code)
}

func TestOverLimit(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(Limit(0, 0))
	router.GET("/test", func(c *gin.Context) {
		c.String(200, "OK")
	})

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, r)

	//assertions here
	assert.Equal(429, w.Code)
}
