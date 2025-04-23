package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"nurcenter/internals/database"
	"nurcenter/internals/models"
)

type TodoRequest struct {
	Title       string `json:"title" binding:"required,max=255"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

func GetTodos(c *gin.Context) {
	userID := c.GetUint("userID")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	isDone := c.Query("is_done")
	offset := (page - 1) * limit

	query := database.DB.Where("user_id = ?", userID)
	if isDone != "" {
		isDoneBool, _ := strconv.ParseBool(isDone)
		query = query.Where("is_done = ?", isDoneBool)
	}

	var todos []models.Todo
	var total int64
	query.Model(&models.Todo{}).Count(&total)
	query.Limit(limit).Offset(offset).Find(&todos)

	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func CreateTodo(c *gin.Context) {
	userID := c.GetUint("userID")

	var req TodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	todo := models.Todo{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		IsDone:      req.IsDone,
	}

	if err := database.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
	userID := c.GetUint("userID")
	id := c.Param("id")

	var todo models.Todo
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	var req TodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	todo.Title = req.Title
	todo.Description = req.Description
	todo.IsDone = req.IsDone

	if err := database.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	userID := c.GetUint("userID")
	id := c.Param("id")

	var todo models.Todo
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := database.DB.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
