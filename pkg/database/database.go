package database

import (
	"github.com/shreyas100100/ExpenseTracker/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	logger = logrus.New()
)

func InitDB() {
	var err error
	dsn := "host=localhost user=postgres password=user dbname=expensetracker port=5432 sslmode=disable TimeZone=Asia/Kolkata"
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
