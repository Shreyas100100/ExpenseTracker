package service

import (
	"errors"
	"sync"
	"time"

	"github.com/shreyas100100/ExpenseTracker/internal/model"
)

var (
	currentID int = 1
	expenses  []model.Expense
	mu        sync.Mutex
)

func CreateExpense(exp model.Expense) model.Expense {
	mu.Lock()
	defer mu.Unlock()

	exp.ID = currentID
	currentID++
	expenses = append(expenses, exp)
	return exp
}

func GetAllExpenses() []model.Expense {
	mu.Lock()
	defer mu.Unlock()
	return append([]model.Expense{}, expenses...)
}

func GetExpenseByID(expenseID int) (model.Expense, error) {
	mu.Lock()
	defer mu.Unlock()

	for _, exp := range expenses {
		if exp.ID == expenseID {
			return exp, nil
		}
	}
	return model.Expense{}, errors.New("expense not found")
}

func GetExpensesByDate(dateString string) ([]model.Expense, error) {
	targetDate, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		return nil, err
	}
	year, month, day := targetDate.Date()

	mu.Lock()
	defer mu.Unlock()

	var filtered []model.Expense
	for _, exp := range expenses {
		expYear, expMonth, expDay := exp.Date.Date()
		if year == expYear && month == expMonth && day == expDay {
			filtered = append(filtered, exp)
		}
	}
	return filtered, nil
}

func UpdateExpense(updatedExp model.Expense) (model.Expense, error) {
	mu.Lock()
	defer mu.Unlock()

	for i, exp := range expenses {
		if exp.ID == updatedExp.ID {
			expenses[i] = updatedExp
			return updatedExp, nil
		}
	}
	return model.Expense{}, errors.New("expense not found")
}

func DeleteExpense(expenseID int) error {
	mu.Lock()
	defer mu.Unlock()

	for i, exp := range expenses {
		if exp.ID == expenseID {
			expenses = append(expenses[:i], expenses[i+1:]...)
			return nil
		}
	}
	return errors.New("expense not found")
}
