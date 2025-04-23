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

func TestCreateReminder(t *testing.T) {
	setupTestDB()
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("userID", uint(1))
		c.Next()
	})
	r.POST("/api/reminders", CreateReminder)

	reqBody, _ := json.Marshal(ReminderRequest{
		Title:    "Test Reminder",
		Message:  "Test Message",
		RemindAt: time.Now().Add(time.Hour),
	})
	req, _ := http.NewRequest("POST", "/api/reminders", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var reminder models.Reminder
	json.Unmarshal(w.Body.Bytes(), &reminder)
	assert.Equal(t, "Test Reminder", reminder.Title)
}
