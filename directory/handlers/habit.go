package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"nurcenter/directory/database"
	"nurcenter/directory/models"
)

type HabitRequest struct {
	Name string `json:"name" binding:"required,max=100"`
}

func CreateHabit(c *gin.Context) {
	userID := c.GetUint("userID")

	var req HabitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	habit := models.Habit{
		UserID: userID,
		Name:   req.Name,
	}

	if err := database.DB.Create(&habit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create habit"})
		return
	}

	c.JSON(http.StatusCreated, habit)
}

func LogHabit(c *gin.Context) {
	userID := c.GetUint("userID")
	habitID := c.Param("habitId")

	var habit models.Habit
	if err := database.DB.Where("id = ? AND user_id = ?", habitID, userID).First(&habit).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Habit not found"})
		return
	}

	log := models.HabitLog{
		HabitID: habit.ID,
		LogDate: time.Now(),
	}

	if err := database.DB.Create(&log).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log habit"})
		return
	}

	c.JSON(http.StatusCreated, log)
}

func GetHabits(c *gin.Context) {
	userID := c.GetUint("userID")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var habits []models.Habit
	var total int64
	query := database.DB.Preload("Logs").Where("user_id = ?", userID)
	query.Model(&models.Habit{}).Count(&total)
	query.Limit(limit).Offset(offset).Find(&habits)

	c.JSON(http.StatusOK, gin.H{
		"habits": habits,
		"total":  total,
		"page":   page,
		"limit":  limit,
	})
}
