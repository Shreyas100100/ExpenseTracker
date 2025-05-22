package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/shreyas100100/ExpenseTracker/internal/model"
	"github.com/shreyas100100/ExpenseTracker/internal/service"
)

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func decodeExpense(r *http.Request) (model.Expense, error) {
	var expense model.Expense
	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		return expense, err
	}
	if expense.Date.IsZero() {
		expense.Date = time.Now()
	}
	return expense, nil
}

func parseIDFromURL(r *http.Request) (int, error) {
	return strconv.Atoi(chi.URLParam(r, "id"))
}

func CreateExpense(w http.ResponseWriter, r *http.Request) {
	expense, err := decodeExpense(r)
	if err != nil {
		http.Error(w, "Invalid request body ", http.StatusBadRequest)
		return
	}
	createdExpense := service.CreateExpense(expense)
	respondJSON(w, http.StatusCreated, createdExpense)
}

func GetAllExpenses(w http.ResponseWriter, r *http.Request) {
	allExpenses := service.GetAllExpenses()
	respondJSON(w, http.StatusOK, allExpenses)
}

func GetExpenseById(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromURL(r)
	if err != nil {
		http.Error(w, "Invalid Id type", http.StatusBadRequest)
		return
	}
	exp, err := service.GetExpenseByID(id)
	if err != nil {
		http.Error(w, "Expense with given id not found", http.StatusNotFound)
		return
	}
	respondJSON(w, http.StatusOK, exp)
}

func GetExpensesByDate(w http.ResponseWriter, r *http.Request) {
	dateStr := chi.URLParam(r, "date")
	expensesByDate, err := service.GetExpensesByDate(dateStr)
	if err != nil {
		http.Error(w, "Invalid Date format.  Use ISO format (2023-06-15T00:00:00Z)", http.StatusBadRequest)
		return
	}
	respondJSON(w, http.StatusOK, expensesByDate)
}

func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromURL(r)
	if err != nil {
		http.Error(w, "Invalid Id type", http.StatusBadRequest)
		return
	}
	var expense model.Expense
	err = json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	expense.ID = id
	updatedExpense, err := service.UpdateExpense(expense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	respondJSON(w, http.StatusOK, updatedExpense)
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromURL(r)
	if err != nil {
		http.Error(w, "Invalid Id type", http.StatusBadRequest)
		return
	}

	err = service.DeleteExpense(id)

	if err != nil {
		http.Error(w, "Expense not found", http.StatusNotFound)
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "Expense deleted successfully"})
}
