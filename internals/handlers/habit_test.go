package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"nurcenter/internals/models"
)

func TestCreateHabit(t *testing.T) {
	setupTestDB()
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("userID", uint(1))
		c.Next()
	})
	r.POST("/api/habits", CreateHabit)

	reqBody, _ := json.Marshal(HabitRequest{
		Name: "Test Habit",
	})
	req, _ := http.NewRequest("POST", "/api/habits", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var habit models.Habit
	json.Unmarshal(w.Body.Bytes(), &habit)
	assert.Equal(t, "Test Habit", habit.Name)
}
