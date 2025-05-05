package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"nurcenter/productivity-service/client"
	"nurcenter/productivity-service/config"
	"nurcenter/productivity-service/models"
	"strconv"
)

func GetHabits(c *gin.Context, db *gorm.DB, cfg *config.Config) {
	userID := c.GetUint("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	if err := client.VerifyUser(userID, cfg); err != nil {
		c.JSON(500, gin.H{"error": "Failed to verify user"})
		return
	}

	var habits []models.Habit
	if err := db.Where("user_id = ?", userID).Offset(offset).Limit(limit).Find(&habits).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to get habits"})
		return
	}

	var total int64
	db.Model(&models.Habit{}).Where("user_id = ?", userID).Count(&total)

	c.JSON(200, gin.H{
		"habits": habits,
		"total":  total,
		"page":   page,
		"limit":  limit,
	})
}

func CreateHabit(c *gin.Context, db *gorm.DB, cfg *config.Config) {
	userID := c.GetUint("user_id")

	var input struct {
		Name      string `json:"name" binding:"required"`
		Frequency string `json:"frequency" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := client.VerifyUser(userID, cfg); err != nil {
		c.JSON(500, gin.H{"error": "Failed to verify user"})
		return
	}

	habit := models.Habit{
		UserID:    userID,
		Name:      input.Name,
		Frequency: input.Frequency,
	}

	if err := db.Create(&habit).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create habit"})
		return
	}

	c.JSON(201, habit)
}
