package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"nurcenter/productivity-service/config"
	"nurcenter/productivity-service/database"
	"nurcenter/productivity-service/handlers"
	"nurcenter/productivity-service/middleware"
)

func main() {
	cfg := config.Load()
	db := database.InitDB(cfg)

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.LoggingMiddleware())

	protected := r.Group("/").Use(middleware.AuthMiddleware())
	{
		protected.GET("/todos", func(c *gin.Context) { handlers.GetTodos(c, db, cfg) })
		protected.POST("/todos", func(c *gin.Context) { handlers.CreateTodo(c, db, cfg) })
		protected.GET("/todos/:id", func(c *gin.Context) { handlers.GetTodoByID(c, db) })
		protected.PUT("/todos/:id", func(c *gin.Context) { handlers.UpdateTodo(c, db) })
		protected.DELETE("/todos/:id", func(c *gin.Context) { handlers.DeleteTodo(c, db) })
		protected.GET("/todos/filter", func(c *gin.Context) { handlers.FilterTodos(c, db) })
		protected.GET("/todos/search", func(c *gin.Context) { handlers.SearchTodos(c, db) })

		protected.GET("/notes", func(c *gin.Context) { handlers.GetNotes(c, db, cfg) })
		protected.POST("/notes", func(c *gin.Context) { handlers.CreateNote(c, db, cfg) })
		protected.GET("/notes/:id", func(c *gin.Context) { handlers.GetNoteByID(c, db) })
		protected.PUT("/notes/:id", func(c *gin.Context) { handlers.UpdateNote(c, db) })
		protected.DELETE("/notes/:id", func(c *gin.Context) { handlers.DeleteNote(c, db) })
		protected.GET("/notes/filter", func(c *gin.Context) { handlers.FilterNotes(c, db) })
		protected.GET("/notes/search", func(c *gin.Context) { handlers.SearchNotes(c, db) })

		protected.GET("/habits", func(c *gin.Context) { handlers.GetHabits(c, db, cfg) })
		protected.POST("/habits", func(c *gin.Context) { handlers.CreateHabit(c, db, cfg) })
		protected.GET("/goals", func(c *gin.Context) { handlers.GetGoals(c, db, cfg) })
		protected.POST("/finances", func(c *gin.Context) { handlers.CreateFinance(c, db, cfg) })
		protected.POST("/reminders", func(c *gin.Context) { handlers.CreateReminder(c, db, cfg) })
		protected.GET("/dashboard", func(c *gin.Context) { handlers.GetDashboard(c, db) })
	}

	r.Run(":8080")

}
