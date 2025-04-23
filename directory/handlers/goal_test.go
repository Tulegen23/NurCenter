package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"nurcenter/directory/models"
)

func TestCreateGoal(t *testing.T) {
	setupTestDB()
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("userID", uint(1))
		c.Next()
	})
	r.POST("/api/goals", CreateGoal)

	reqBody, _ := json.Marshal(GoalRequest{
		Title:       "Test Goal",
		Description: "Test Description",
		Deadline:    time.Now().AddDate(0, 1, 0),
	})
	req, _ := http.NewRequest("POST", "/api/goals", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var goal models.Goal
	json.Unmarshal(w.Body.Bytes(), &goal)
	assert.Equal(t, "Test Goal", goal.Title)
}
