package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"nurcenter/productivity-service/config"
	"nurcenter/productivity-service/handlers"
	"nurcenter/productivity-service/models"
)

func setupProductivityTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&models.Todo{}, &models.Note{}, &models.Habit{}, &models.Goal{}, &models.Finance{}, &models.Reminder{})
	return db
}

func TestCreateTodo(t *testing.T) {
	db := setupProductivityTestDB()
	cfg := &config.Config{UserServiceURL: "http://localhost:8081"}
	router := gin.Default()
	router.POST("/todos", func(c *gin.Context) {
		c.Set("user_id", uint(1))
		handlers.CreateTodo(c, db, cfg)
	})

	payload := map[string]string{
		"title":       "Test Todo",
		"description": "Test Description",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 201 {
		t.Errorf("Expected status 201, got %d", w.Code)
	}
}

func TestGetTodos(t *testing.T) {
	db := setupProductivityTestDB()
	cfg := &config.Config{UserServiceURL: "http://localhost:8081"}
	router := gin.Default()
	router.GET("/todos", func(c *gin.Context) {
		c.Set("user_id", uint(1))
		handlers.GetTodos(c, db, cfg)
	})

	db.Create(&models.Todo{UserID: 1, Title: "Test Todo"})

	req, _ := http.NewRequest("GET", "/todos?page=1&limit=10", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestCreateNote(t *testing.T) {
	db := setupProductivityTestDB()
	cfg := &config.Config{UserServiceURL: "http://localhost:8081"}
	router := gin.Default()
	router.POST("/notes", func(c *gin.Context) {
		c.Set("user_id", uint(1))
		handlers.CreateNote(c, db, cfg)
	})

	payload := map[string]string{
		"title":    "Test Note",
		"content":  "Test Content",
		"category": "personal",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/notes", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 201 {
		t.Errorf("Expected status 201, got %d", w.Code)
	}
}

func TestCreateHabit(t *testing.T) {
	db := setupProductivityTestDB()
	cfg := &config.Config{UserServiceURL: "http://localhost:8081"}
	router := gin.Default()
	router.POST("/habits", func(c *gin.Context) {
		c.Set("user_id", uint(1))
		handlers.CreateHabit(c, db, cfg)
	})

	payload := map[string]string{
		"name":      "Test Habit",
		"frequency": "daily",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/habits", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 201 {
		t.Errorf("Expected status 201, got %d", w.Code)
	}
}

func TestCreateGoal(t *testing.T) {
	db := setupProductivityTestDB()
	cfg := &config.Config{UserServiceURL: "http://localhost:8081"}
	router := gin.Default()
	router.POST("/goals", func(c *gin.Context) {
		c.Set("user_id", uint(1))
		handlers.CreateGoal(c, db, cfg)
	})

	payload := map[string]interface{}{
		"title":    "Test Goal",
		"deadline": "2025-12-31T23:59:59Z",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/goals", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 201 {
		t.Errorf("Expected status 201, got %d", w.Code)
	}
}

func TestCreateFinance(t *testing.T) {
	db := setupProductivityTestDB()
	cfg := &config.Config{UserServiceURL: "http://localhost:8081"}
	router := gin.Default()
	router.POST("/finances", func(c *gin.Context) {
		c.Set("user_id", uint(1))
		handlers.CreateFinance(c, db, cfg)
	})

	payload := map[string]interface{}{
		"amount":      100.50,
		"description": "Test Finance",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/finances", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 201 {
		t.Errorf("Expected status 201, got %d", w.Code)
	}
}

func TestCreateReminder(t *testing.T) {
	db := setupProductivityTestDB()
	cfg := &config.Config{UserServiceURL: "http://localhost:8081"}
	router := gin.Default()
	router.POST("/reminders", func(c *gin.Context) {
		c.Set("user_id", uint(1))
		handlers.CreateReminder(c, db, cfg)
	})

	payload := map[string]interface{}{
		"title":    "Test Reminder",
		"due_date": "2025-05-10T10:00:00Z",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/reminders", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 201 {
		t.Errorf("Expected status 201, got %d", w.Code)
	}
}
