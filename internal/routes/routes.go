package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/shreyas100100/ExpenseTracker/internal/handler"
)

func SetupRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/api/expenses", handler.GetAllExpenses)
	router.Get("/api/expenses/{id}", handler.GetExpenseById)
	router.Get("/api/expenses/date/{date}", handler.GetExpensesByDate)
	router.Post("/api/expenses", handler.CreateExpense)
	router.Put("/api/expenses/{id}", handler.UpdateExpense)
	router.Delete("/api/expenses/{id}", handler.DeleteExpense)
	return router
}
