package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"nurcenter/user-service/handlers"
	"nurcenter/user-service/models"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	return db
}

func TestRegister(t *testing.T) {
	db := setupTestDB()
	router := gin.Default()
	router.POST("/register", func(c *gin.Context) {
		handlers.Register(c, db)
	})

	payload := map[string]string{
		"username": "testuser",
		"email":    "test@example.com",
		"password": "password123",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 201 {
		t.Errorf("Expected status 201, got %d", w.Code)
	}
}

func TestLogin(t *testing.T) {
	db := setupTestDB()
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		handlers.Login(c, db)
	})

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	db.Create(&models.User{Username: "testuser", Email: "test@example.com", Password: string(hashedPassword)})

	payload := map[string]string{
		"email":    "test@example.com",
		"password": "password123",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestGetUserByID(t *testing.T) {
	db := setupTestDB()
	router := gin.Default()
	router.GET("/users/:id", func(c *gin.Context) {
		handlers.GetUserByID(c, db)
	})

	db.Create(&models.User{ID: 1, Username: "testuser", Email: "test@example.com"})

	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
