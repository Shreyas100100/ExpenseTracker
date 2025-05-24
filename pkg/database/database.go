package database

import (
	"fmt"

	"github.com/shreyas100100/ExpenseTracker/internal/model"
	"github.com/shreyas100100/ExpenseTracker/pkg/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	logger = logrus.New()
)

func InitDB() {
	cfg := config.LoadConfig()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.TimeZone,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.WithError(err).Fatal("Failed to connect to database")
		return
	}
	err = DB.AutoMigrate(&model.Expense{})
	if err != nil {
		logger.WithError(err).Fatal("Failed to migrate database")
		return
	}
	logger.Info("Database connection successful")
}
