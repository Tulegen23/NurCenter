package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"nurcenter/productivity-service/client"
	"nurcenter/productivity-service/config"
	"nurcenter/productivity-service/models"
	"strconv"
)

func GetNotes(c *gin.Context, db *gorm.DB, cfg *config.Config) {
	userID := c.GetUint("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	if err := client.VerifyUser(userID, cfg); err != nil {
		c.JSON(500, gin.H{"error": "Failed to verify user"})
		return
	}

	var notes []models.Note
	if err := db.Where("user_id = ?", userID).Offset(offset).Limit(limit).Find(&notes).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to get notes"})
		return
	}

	var total int64
	db.Model(&models.Note{}).Where("user_id = ?", userID).Count(&total)

	c.JSON(200, gin.H{
		"notes": notes,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func CreateNote(c *gin.Context, db *gorm.DB, cfg *config.Config) {
	userID := c.GetUint("user_id")
	var input struct {
		Title    string `json:"title" binding:"required"`
		Content  string `json:"content"`
		Category string `json:"category"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := client.VerifyUser(userID, cfg); err != nil {
		c.JSON(500, gin.H{"error": "Failed to verify user"})
		return
	}

	note := models.Note{
		UserID:   userID,
		Title:    input.Title,
		Content:  input.Content,
		Category: input.Category,
	}

	if err := db.Create(&note).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create note"})
		return
	}

	c.JSON(201, note)
}

func GetNoteByID(c *gin.Context, db *gorm.DB) {
	userID := c.GetUint("user_id")
	id, _ := strconv.Atoi(c.Param("id"))

	var note models.Note
	if err := db.Where("id = ? AND user_id = ?", id, userID).First(&note).Error; err != nil {
		c.JSON(404, gin.H{"error": "Note not found"})
		return
	}

	c.JSON(200, note)
}

func UpdateNote(c *gin.Context, db *gorm.DB) {
	userID := c.GetUint("user_id")
	id, _ := strconv.Atoi(c.Param("id"))

	var note models.Note
	if err := db.Where("id = ? AND user_id = ?", id, userID).First(&note).Error; err != nil {
		c.JSON(404, gin.H{"error": "Note not found"})
		return
	}

	var input struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		Category string `json:"category"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Model(&note).Updates(models.Note{
		Title:    input.Title,
		Content:  input.Content,
		Category: input.Category,
	})

	c.JSON(200, note)
}

func DeleteNote(c *gin.Context, db *gorm.DB) {
	userID := c.GetUint("user_id")
	id, _ := strconv.Atoi(c.Param("id"))

	var note models.Note
	if err := db.Where("id = ? AND user_id = ?", id, userID).First(&note).Error; err != nil {
		c.JSON(404, gin.H{"error": "Note not found"})
		return
	}

	db.Delete(&note)
	c.JSON(204, nil)
}

func FilterNotes(c *gin.Context, db *gorm.DB) {
	userID := c.GetUint("user_id")
	category := c.Query("category")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var notes []models.Note
	query := db.Where("user_id = ?", userID)
	if category != "" {
		query = query.Where("category = ?", category)
	}

	if err := query.Offset(offset).Limit(limit).Find(&notes).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to filter notes"})
		return
	}

	var total int64
	query.Count(&total)

	c.JSON(200, gin.H{
		"notes": notes,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func SearchNotes(c *gin.Context, db *gorm.DB) {
	userID := c.GetUint("user_id")
	title := c.Query("title")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var notes []models.Note
	query := db.Where("user_id = ? AND title ILIKE ?", userID, "%"+title+"%")
	if err := query.Offset(offset).Limit(limit).Find(&notes).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to search notes"})
		return
	}

	var total int64
	query.Count(&total)

	c.JSON(200, gin.H{
		"notes": notes,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}
