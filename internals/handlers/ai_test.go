package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/gorm"

	_ "nurcenter/internals/database"
)

func TestAnalyzeNotes(t *testing.T) {
	setupTestDB()
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("userID", uint(1))
		c.Next()
	})
	r.POST("/api/ai/analyze", AnalyzeNotes)

	req, _ := http.NewRequest("POST", "/api/ai/analyze", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Contains(t, response["analysis"], "Summary")
}
