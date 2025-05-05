package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"nurcenter/productivity-service/config"
)

func InitDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.ProductivityDBHost,
		cfg.ProductivityDBPort,
		cfg.ProductivityDBUser,
		cfg.ProductivityDBPassword,
		cfg.ProductivityDBName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to productivity database")
	}
	return db
}
