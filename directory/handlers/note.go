package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"nurcenter/directory/database"
	"nurcenter/directory/models"
)

type NoteRequest struct {
	Title    string `json:"title" binding:"required,max=200"`
	Content  string `json:"content"`
	Category string `json:"category" binding:"max=50"`
}

func CreateNote(c *gin.Context) {
	userID := c.GetUint("userID")

	var req NoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	note := models.Note{
		UserID:   userID,
		Title:    req.Title,
		Content:  req.Content,
		Category: req.Category,
	}

	if err := database.DB.Create(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}

	c.JSON(http.StatusCreated, note)
}

func GetNotes(c *gin.Context) {
	userID := c.GetUint("userID")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	category := c.Query("category")
	offset := (page - 1) * limit

	query := database.DB.Where("user_id = ?", userID)
	if category != "" {
		query = query.Where("category = ?", category)
	}

	var notes []models.Note
	var total int64
	query.Model(&models.Note{}).Count(&total)
	query.Limit(limit).Offset(offset).Find(&notes)

	c.JSON(http.StatusOK, gin.H{
		"notes": notes,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func UpdateNote(c *gin.Context) {
	userID := c.GetUint("userID")
	id := c.Param("id")

	var note models.Note
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	var req NoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	note.Title = req.Title
	note.Content = req.Content
	note.Category = req.Category

	if err := database.DB.Save(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
		return
	}

	c.JSON(http.StatusOK, note)
}

func DeleteNote(c *gin.Context) {
	userID := c.GetUint("userID")
	id := c.Param("id")

	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Note{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
}
