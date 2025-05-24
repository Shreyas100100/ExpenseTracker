package service

import (
	"log"

	"github.com/shreyas100100/ExpenseTracker/internal/model"
	"github.com/shreyas100100/ExpenseTracker/internal/repository"
)

func CreateExpense(exp model.Expense) model.Expense {
	createdExp, err := repository.CreateExpense(exp)
	if err != nil {
		log.Fatalf("Error while creating an expense %f", err)
	}
	return createdExp
}

func GetAllExpenses() []model.Expense {
	expenses, err := repository.GetAllExpenses()
	if err != nil {

		log.Fatalf("Error while fetching all expenses %f", err)
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
