package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"nurcenter/directory/database"
	"nurcenter/directory/models"
)

type GoalRequest struct {
	Title       string    `json:"title" binding:"required,max=200"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
}

type SubgoalRequest struct {
	Title  string `json:"title" binding:"required,max=200"`
	IsDone bool   `json:"is_done"`
}

func CreateGoal(c *gin.Context) {
	userID := c.GetUint("userID")

	var req GoalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	goal := models.Goal{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Deadline:    req.Deadline,
	}

	if err := database.DB.Create(&goal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create goal"})
		return
	}

	c.JSON(http.StatusCreated, goal)
}

func GetGoals(c *gin.Context) {
	userID := c.GetUint("userID")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	isCompleted := c.Query("is_completed")
	offset := (page - 1) * limit

	query := database.DB.Preload("Subgoals").Where("user_id = ?", userID)
	if isCompleted != "" {
		isCompletedBool, _ := strconv.ParseBool(isCompleted)
		query = query.Where("is_completed = ?", isCompletedBool)
	}

	var goals []models.Goal
	var total int64
	query.Model(&models.Goal{}).Count(&total)
	query.Limit(limit).Offset(offset).Find(&goals)

	c.JSON(http.StatusOK, gin.H{
		"goals": goals,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func AddSubgoal(c *gin.Context) {
	userID := c.GetUint("userID")
	goalID := c.Param("goalId")

	var goal models.Goal
	if err := database.DB.Where("id = ? AND user_id = ?", goalID, userID).First(&goal).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Goal not found"})
		return
	}

	var req SubgoalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	subgoal := models.Subgoal{
		GoalID: goal.ID,
		Title:  req.Title,
		IsDone: req.IsDone,
	}

	if err := database.DB.Create(&subgoal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subgoal"})
		return
	}

	c.JSON(http.StatusCreated, subgoal)
}

func UpdateSubgoalStatus(c *gin.Context) {
	userID := c.GetUint("userID")
	subgoalID := c.Param("subgoalId")

	var subgoal models.Subgoal
	if err := database.DB.
		Joins("JOIN goals ON goals.id = subgoals.goal_id").
		Where("subgoals.id = ? AND goals.user_id = ?", subgoalID, userID).
		First(&subgoal).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subgoal not found"})
		return
	}

	var req SubgoalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	subgoal.Title = req.Title
	subgoal.IsDone = req.IsDone

	if err := database.DB.Save(&subgoal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update subgoal"})
		return
	}

	c.JSON(http.StatusOK, subgoal)
}
