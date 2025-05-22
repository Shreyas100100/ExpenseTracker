package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/shreyas100100/ExpenseTracker/internal/handler"
)

func main() {

	router := chi.NewRouter()
	router.Get("/api/expenses", handler.GetAllExpenses)
	router.Get("/api/expenses/{id}", handler.GetExpenseById)
	router.Get("/api/expenses/date/{date}", handler.GetExpensesByDate)
	router.Post("/api/expenses", handler.CreateExpense)
	router.Put("/api/expenses/{id}", handler.UpdateExpense)
	router.Delete("/api/expenses/{id}", handler.DeleteExpense)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
