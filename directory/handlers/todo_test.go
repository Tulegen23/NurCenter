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

func TestCreateTodo(t *testing.T) {
	setupTestDB()
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Mock middleware
	r.Use(func(c *gin.Context) {
		c.Set("userID", uint(1))
		c.Next()
	})
	r.POST("/api/todos", CreateTodo)

	reqBody, _ := json.Marshal(TodoRequest{
		Title:       "Test Todo",
		Description: "Test Description",
		IsDone:      false,
	})
	req, _ := http.NewRequest("POST", "/api/todos", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var todo models.Todo
	json.Unmarshal(w.Body.Bytes(), &todo)
	assert.Equal(t, "Test Todo", todo.Title)
}
