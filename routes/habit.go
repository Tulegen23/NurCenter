package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nurcenter/database"
	"nurcenter/models"
)

func RegisterHabitRoutes(r *gin.Engine) {
	h := r.Group("/habits")

	h.GET("/", func(c *gin.Context) {
		var habits []models.Habit
		database.DB.Find(&habits)
		c.JSON(http.StatusOK, habits)
	})

	h.POST("/", func(c *gin.Context) {
		var habit models.Habit
		if err := c.ShouldBindJSON(&habit); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		database.DB.Create(&habit)
		c.JSON(http.StatusOK, habit)
	})
}
