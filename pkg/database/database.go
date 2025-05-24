package database

import (
	"log"

	"github.com/shreyas100100/ExpenseTracker/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "host=localhost user=postgres password=user dbname=expensetracker port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database ", err)
		return
	}
	err = DB.AutoMigrate(&model.Expense{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return
	}
	log.Println("Database connection success")

}
