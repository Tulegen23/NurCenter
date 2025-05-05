package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-resty/resty/v2"
	"gorm.io/gorm"
	"nurcenter/productivity-service/client"
	"nurcenter/productivity-service/config"
	"nurcenter/productivity-service/models"
	"strconv"
)

func GetTodos(c *gin.Context, db *gorm.DB, cfg *config.Config) {
	userID := c.GetUint("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	if err := client.VerifyUser(userID, cfg); err != nil {
		c.JSON(500, gin.H{"error": "Failed to verify user"})
		return
	}

	var todos []models.Todo
	if err := db.Where("user_id = ?", userID).Offset(offset).Limit(limit).Find(&todos).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to get todos"})
		return
	}

	var total int64
	db.Model(&models.Todo{}).Where("user_id = ?", userID).Count(&total)

	c.JSON(200, gin.H{
		"todos": todos,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func CreateTodo(c *gin.Context, db *gorm.DB, cfg *config.Config) {
	userID := c.GetUint("user_id")
	var input struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := client.VerifyUser(userID, cfg); err != nil {
		c.JSON(500, gin.H{"error": "Failed to verify user"})
		return
	}

	todo := models.Todo{
		UserID:      userID,
		Title:       input.Title,
		Description: input.Description,
	}

	if err := db.Create(&todo).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(201, todo)
}

func GetTodoByID(c *gin.Context, db *gorm.DB) {
	userID := c.GetUint("user_id")
	id, _ := strconv.Atoi(c.Param("id"))

	var todo models.Todo
	if err := db.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error; err != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(200, todo)
}

func UpdateTodo(c *gin.Context, db *gorm.DB) {
	userID := c.GetUint("user_id")
	id, _ := strconv.Atoi(c.Param("id"))

	var todo models.Todo
	if err := db.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error; err != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		IsDone      bool   `json:"is_done"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Model(&todo).Updates(models.Todo{
		Title:       input.Title,
		Description: input.Description,
		IsDone:      input.IsDone,
	})

	c.JSON(200, todo)
}

func DeleteTodo(c *gin.Context, db *gorm.DB) {
	userID := c.GetUint("user_id")
	id, _ := strconv.Atoi(c.Param("id"))

	var todo models.Todo
	if err := db.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error; err != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	db.Delete(&todo)
	c.JSON(204, nil)
}

func FilterTodos(c *gin.Context, db *gorm.DB) {
	userID := c.GetUint("user_id")
	isDone := c.Query("is_done")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var todos []models.Todo
	query := db.Where("user_id = ?", userID)
	if isDone != "" {
		done, _ := strconv.ParseBool(isDone)
		query = query.Where("is_done = ?", done)
	}

	if err := query.Offset(offset).Limit(limit).Find(&todos).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to filter todos"})
		return
	}

	var total int64
	query.Count(&total)

	c.JSON(200, gin.H{
		"todos": todos,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func SearchTodos(c *gin.Context, db *gorm.DB) {
	userID := c.GetUint("user_id")
	title := c.Query("title")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var todos []models.Todo
	query := db.Where("user_id = ? AND title ILIKE ?", userID, "%"+title+"%")
	if err := query.Offset(offset).Limit(limit).Find(&todos).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to search todos"})
		return
	}

	var total int64
	query.Count(&total)

	c.JSON(200, gin.H{
		"todos": todos,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}
