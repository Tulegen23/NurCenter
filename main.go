package main

import (
	"github.com/joho/godotenv"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"nurcenter/internals/database"
	"nurcenter/internals/handlers"
	"nurcenter/internals/middleware"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	// Initialize database
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Create Gin application
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup routes
	setupRoutes(r)

	// Start server
	port := ":8080"
	log.Printf("Starting server on port %s", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRoutes(r *gin.Engine) {
	// Public routes
	public := r.Group("/api")
	{
		public.POST("/register", handlers.Register)
		public.POST("/login", handlers.Login)
	}

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// Todos
		protected.GET("/todos", handlers.GetTodos)
		protected.POST("/todos", handlers.CreateTodo)
		protected.PUT("/todos/:id", handlers.UpdateTodo)
		protected.DELETE("/todos/:id", handlers.DeleteTodo)

		// Notes
		protected.GET("/notes", handlers.GetNotes)
		protected.POST("/notes", handlers.CreateNote)
		protected.PUT("/notes/:id", handlers.UpdateNote)
		protected.DELETE("/notes/:id", handlers.DeleteNote)

		// Habits
		protected.GET("/habits", handlers.GetHabits)
		protected.POST("/habits", handlers.CreateHabit)
		protected.POST("/habits/:habitId/log", handlers.LogHabit)

		// Goals
		protected.GET("/goals", handlers.GetGoals)
		protected.POST("/goals", handlers.CreateGoal)
		protected.POST("/goals/:goalId/subgoals", handlers.AddSubgoal)
		protected.PUT("/subgoals/:subgoalId", handlers.UpdateSubgoalStatus)

		//// Finances
		//protected.GET("/finances", handlers.GetFinances)
		//protected.POST("/finances", handlers.AddFinance)

		// Reminders
		protected.GET("/reminders", handlers.GetReminders)
		protected.POST("/reminders", handlers.CreateReminder)
		protected.DELETE("/reminders/:id", handlers.DeleteReminder)

		// Dashboard
		protected.GET("/dashboard", handlers.GetDashboard)

		// AI
		protected.POST("/ai/analyze", handlers.AnalyzeNotes)
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
			"time":   time.Now().Format(time.RFC3339),
		})
	})
}
