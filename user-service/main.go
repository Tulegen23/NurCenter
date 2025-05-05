package main

import (
	"github.com/gin-gonic/gin"
	"nurcenter/user-service/config"
	"nurcenter/user-service/database"
	"nurcenter/user-service/handlers"
	"nurcenter/user-service/middleware"
)

func main() {
	cfg := config.Load()
	db := database.InitDB(cfg)

	r := gin.Default()
	r.Use(middleware.LoggingMiddleware())

	r.POST("/register", func(c *gin.Context) { handlers.Register(c, db) })
	r.POST("/login", func(c *gin.Context) { handlers.Login(c, db) })
	r.GET("/users/:id", func(c *gin.Context) { handlers.GetUserByID(c, db) })

	r.Run(":8081")

}
