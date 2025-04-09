package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nurcenter/database"
	"nurcenter/models"
)

func RegisterNoteRoutes(r *gin.Engine) {
	n := r.Group("/notes")

	n.GET("/", func(c *gin.Context) {
		var notes []models.Note
		database.DB.Find(&notes)
		c.JSON(http.StatusOK, notes)
	})

	n.POST("/", func(c *gin.Context) {
		var note models.Note
		if err := c.ShouldBindJSON(&note); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		database.DB.Create(&note)
		c.JSON(http.StatusOK, note)
	})
}
