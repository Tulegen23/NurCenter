package database

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"nurcenter/user-service/config"
)

func InitDB(cfg *config.Config) *gorm.DB {
	// Construct the DSN for GORM
	gormDSN := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.UserDBHost,
		cfg.UserDBPort,
		cfg.UserDBUser,
		cfg.UserDBPassword,
		cfg.UserDBName,
	)

	db, err := gorm.Open(postgres.Open(gormDSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect to user database")
	}

	// Construct the DSN with scheme for golang-migrate
	migrateDSN := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.UserDBUser,
		cfg.UserDBPassword,
		cfg.UserDBHost,
		cfg.UserDBPort,
		cfg.UserDBName,
	)

	// Create a new migrate instance with file-based migrations from nurcenter/migrations
	m, err := migrate.New(
		"file://../migrations", // Path to the migrations directory
		migrateDSN,             // Use the DSN with scheme
	)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize migrate instance: %v", err))
	}

	// Apply migrations (up)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(fmt.Sprintf("failed to apply migrations: %v", err))
	}

	fmt.Println("Database migrations applied successfully")
	return db
}
