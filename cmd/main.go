package main

import (
	"github.com/gin-gonic/gin"
	"nurcenter/database"
	"nurcenter/routes"
)

func main() {
	r := gin.Default()
	database.Connect()

	routes.RegisterToDoRoutes(r)
	routes.RegisterPomodoroRoutes(r)
	routes.RegisterNoteRoutes(r)
	routes.RegisterHabitRoutes(r)

	r.Run(":8080")
}
