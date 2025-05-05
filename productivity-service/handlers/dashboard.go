package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"nurcenter/productivity-service/models"
)

func GetDashboard(c *gin.Context, db *gorm.DB) {
	userID := c.GetUint("user_id")

	var todoCount int64
	db.Model(&models.Todo{}).Where("user_id = ? AND is_done = ?", userID, true).Count(&todoCount)

	var noteCount int64
	db.Model(&models.Note{}).Where("user_id = ?", userID).Count(&noteCount)

	var habitCount int64
	db.Model(&models.Habit{}).Where("user_id = ?", userID).Count(&habitCount)

	var goalCount int64
	db.Model(&models.Goal{}).Where("user_id = ?", userID).Count(&goalCount)

	c.JSON(200, gin.H{
		"completed_todos": todoCount,
		"total_notes":     noteCount,
		"total_habits":    habitCount,
		"total_goals":     goalCount,
	})
}
