package repository

import (
	"errors"
	"time"

	"github.com/shreyas100100/ExpenseTracker/internal/model"
	"github.com/shreyas100100/ExpenseTracker/pkg/database"
)

func CreateExpense(exp model.Expense) (model.Expense, error) {
	result := database.DB.Create(&exp)
	return exp, result.Error
}

func GetAllExpenses() ([]model.Expense, error) {
	var expenses []model.Expense
	result := database.DB.Find(&expenses)
	return expenses, result.Error
}

func GetExpenseByID(id int) (model.Expense, error) {
	var expense model.Expense
	result := database.DB.First(&expense, id)
	if result.Error != nil {
		return model.Expense{}, result.Error
	}
	return expense, nil
}

func GetExpensesByDate(dateStr string) ([]model.Expense, error) {

	targetDate, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return nil, err
	}
	dateOnly := targetDate.Format("2006-01-02")
	var expenses []model.Expense
	result := database.DB.Where("DATE(date) = ?", dateOnly).Find(&expenses)
	return expenses, result.Error
}

func UpdateExpense(updatedExp model.Expense) (model.Expense, error) {
	var existing model.Expense
	if err := database.DB.First(&existing, updatedExp.ID).Error; err != nil {
		return model.Expense{}, errors.New("expense not found")
	}
	result := database.DB.Save(&updatedExp)
	return updatedExp, result.Error
}

func DeleteExpense(id int) error {
	result := database.DB.Delete(&model.Expense{}, id)
	if result.RowsAffected == 0 {
		return errors.New("expense not found")
	}
	return result.Error
}
