package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"nurcenter/directory/database"
	"nurcenter/directory/models"
)

func AnalyzeNotes(c *gin.Context) {
	userID := c.GetUint("userID")

	var notes []models.Note
	if err := database.DB.Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes"})
		return
	}

	// Mock AI analysis (replace with real AI integration, e.g., OpenAI)
	var noteContents []string
	for _, note := range notes {
		noteContents = append(noteContents, note.Content)
	}
	analysis := "Summary: " + strings.Join(noteContents, " ")

	c.JSON(http.StatusOK, gin.H{
		"analysis": analysis,
	})
}
