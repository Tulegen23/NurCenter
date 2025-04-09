package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"nurcenter/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=nurbol23 dbname=nurcenter port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	db.AutoMigrate(&models.ToDo{}, &models.Pomodoro{}, &models.Note{}, &models.Habit{})

	DB = db
}
