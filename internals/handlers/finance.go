package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"

	"nurcenter/internals/database"
	"nurcenter/internals/models"
)

type FinanceRequest struct {
	Amount      float64 `json:"amount" binding:"required,gt=0"`
	Type        string  `json:"type" binding:"required,oneof=income expense"`
	Category    string  `json:"category" binding:"max=100"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
}

func AddFinance(c *gin.Context) {
	userID := c.GetUint("userID")

	var req FinanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		date = time.Now()
	}

	finance := models.Finance{
		UserID:      userID,
		Amount:      decimal.NewFromFloat(req.Amount),
		Type:        req.Type,
		Category:    req.Category,
		Description: req.Description,
		Date:        date,
	}

	if err := database.DB.Create(&finance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create finance"})
		return
	}

	c.JSON(http.StatusCreated, finance)
}

func GetFinances(c *gin.Context) {
	userID := c.GetUint("userID")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	typeFilter := c.Query("type")
	category := c.Query("category")
	offset := (page - 1) * limit

	query := database.DB.Where("user_id = ?", userID)
	if typeFilter != "" {
		query = query.Where("type = ?", typeFilter)
	}
	if category != "" {
		query = query.Where("category = ?", category)
	}

	var finances []models.Finance
	var total int64
	query.Model(&models.Finance{}).Count(&total)
	query.Limit(limit).Offset(offset).Find(&finances)

	c.JSON(http.StatusOK, gin.H{
		"finances": finances,
		"total":    total,
		"page":     page,
		"limit":    limit,
	})
}
