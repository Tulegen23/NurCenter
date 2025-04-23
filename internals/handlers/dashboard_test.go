package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetDashboard(t *testing.T) {
	setupTestDB()
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("userID", uint(1))
		c.Next()
	})
	r.GET("/api/dashboard", GetDashboard)

	req, _ := http.NewRequest("GET", "/api/dashboard", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response DashboardResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, int64(0), response.TotalTodos)
}
