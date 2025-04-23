package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"nurcenter/directory/database"
	"nurcenter/directory/models"
)

type ReminderRequest struct {
	Title    string    `json:"title" binding:"required,max=255"`
	Message  string    `json:"message"`
	RemindAt time.Time `json:"remind_at" binding:"required"`
}

func CreateReminder(c *gin.Context) {
	userID := c.GetUint("userID")

	var req ReminderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	reminder := models.Reminder{
		UserID:   userID,
		Title:    req.Title,
		Message:  req.Message,
		RemindAt: req.RemindAt,
	}

	if err := database.DB.Create(&reminder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reminder"})
		return
	}

	c.JSON(http.StatusCreated, reminder)
}

func GetReminders(c *gin.Context) {
	userID := c.GetUint("userID")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	isSent := c.Query("is_sent")
	offset := (page - 1) * limit

	query := database.DB.Where("user_id = ? AND remind_at > ?", userID, time.Now())
	if isSent != "" {
		isSentBool, _ := strconv.ParseBool(isSent)
		query = query.Where("is_sent = ?", isSentBool)
	}

	var reminders []models.Reminder
	var total int64
	query.Model(&models.Reminder{}).Count(&total)
	query.Limit(limit).Offset(offset).Find(&reminders)

	c.JSON(http.StatusOK, gin.H{
		"reminders": reminders,
		"total":     total,
		"page":      page,
		"limit":     limit,
	})
}

func DeleteReminder(c *gin.Context) {
	userID := c.GetUint("userID")
	id := c.Param("id")

	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Reminder{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reminder not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reminder deleted"})
}
