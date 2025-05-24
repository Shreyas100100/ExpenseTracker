package service

import (
	"github.com/shreyas100100/ExpenseTracker/internal/model"
	"github.com/shreyas100100/ExpenseTracker/internal/repository"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func CreateExpense(exp model.Expense) model.Expense {
	createdExp, err := repository.CreateExpense(exp)
	if err != nil {
		logger.WithError(err).Error("CreateExpense: error creating expense")
	}
	return createdExp
}

func GetAllExpenses() []model.Expense {
	expenses, err := repository.GetAllExpenses()
	if err != nil {
		logger.WithError(err).Error("GetAllExpenses: error fetching expenses")
	}
	return expenses
}

func GetExpenseByID(id int) (model.Expense, error) {
	return repository.GetExpenseByID(id)
}

func GetExpensesByDate(dateStr string) ([]model.Expense, error) {
	return repository.GetExpensesByDate(dateStr)
}

func UpdateExpense(exp model.Expense) (model.Expense, error) {
	return repository.UpdateExpense(exp)
}

func DeleteExpense(id int) error {
	return repository.DeleteExpense(id)
}
