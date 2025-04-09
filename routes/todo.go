package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nurcenter/database"
	"nurcenter/models"
)

func RegisterToDoRoutes(r *gin.Engine) {
	todo := r.Group("/todos")

	todo.GET("/", func(c *gin.Context) {
		var todos []models.ToDo
		database.DB.Find(&todos)
		c.JSON(http.StatusOK, todos)
	})

	todo.POST("/", func(c *gin.Context) {
		var todo models.ToDo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		database.DB.Create(&todo)
		c.JSON(http.StatusOK, todo)
	})

	todo.PUT("/:id", func(c *gin.Context) {
		var todo models.ToDo
		id := c.Param("id")
		if err := database.DB.First(&todo, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Не найдено"})
			return
		}
		c.ShouldBindJSON(&todo)
		database.DB.Save(&todo)
		c.JSON(http.StatusOK, todo)
	})

	todo.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		database.DB.Delete(&models.ToDo{}, id)
		c.JSON(http.StatusOK, gin.H{"message": "Удалено"})
	})
}
