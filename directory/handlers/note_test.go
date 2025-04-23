package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"nurcenter/directory/models"
)

func TestCreateNote(t *testing.T) {
	setupTestDB()
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("userID", uint(1))
		c.Next()
	})
	r.POST("/api/notes", CreateNote)

	reqBody, _ := json.Marshal(NoteRequest{
		Title:    "Test Note",
		Content:  "Test Content",
		Category: "Personal",
	})
	req, _ := http.NewRequest("POST", "/api/notes", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var note models.Note
	json.Unmarshal(w.Body.Bytes(), &note)
	assert.Equal(t, "Test Note", note.Title)
}
