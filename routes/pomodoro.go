package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nurcenter/database"
	"nurcenter/models"
)

func RegisterPomodoroRoutes(r *gin.Engine) {
	p := r.Group("/pomodoros")

	p.GET("/", func(c *gin.Context) {
		var data []models.Pomodoro
		database.DB.Find(&data)
		c.JSON(http.StatusOK, data)
	})

	p.POST("/", func(c *gin.Context) {
		var session models.Pomodoro
		if err := c.ShouldBindJSON(&session); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		database.DB.Create(&session)
		c.JSON(http.StatusOK, session)
	})
}
