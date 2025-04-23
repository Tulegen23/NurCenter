package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"nurcenter/internals/database"
	"nurcenter/internals/models"
)

type DashboardResponse struct {
	TotalTodos     int64   `json:"total_todos"`
	CompletedTodos int64   `json:"completed_todos"`
	ActiveHabits   int64   `json:"active_habits"`
	TotalIncome    float64 `json:"total_income"`
	TotalExpenses  float64 `json:"total_expenses"`
}

func GetDashboard(c *gin.Context) {
	userID := c.GetUint("userID")

	var response DashboardResponse

	// Aggregate todos
	type TodoStats struct {
		TotalTodos     int64
		CompletedTodos int64
	}
	var todoStats TodoStats
	database.DB.Model(&models.Todo{}).
		Select("COUNT(*) as total_todos, SUM(CASE WHEN is_done THEN 1 ELSE 0 END) as completed_todos").
		Where("user_id = ?", userID).
		Scan(&todoStats)
	response.TotalTodos = todoStats.TotalTodos
	response.CompletedTodos = todoStats.CompletedTodos

	// Habits
	database.DB.Model(&models.Habit{}).
		Where("user_id = ?", userID).
		Count(&response.ActiveHabits)

	// Finances
	type FinanceStats struct {
		TotalIncome   float64
		TotalExpenses float64
	}
	var financeStats FinanceStats
	database.DB.Model(&models.Finance{}).
		Select("SUM(CASE WHEN type = 'income' THEN amount ELSE 0 END) as total_income, SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END) as total_expenses").
		Where("user_id = ?", userID).
		Scan(&financeStats)
	response.TotalIncome = financeStats.TotalIncome
	response.TotalExpenses = financeStats.TotalExpenses

	c.JSON(http.StatusOK, response)
}
