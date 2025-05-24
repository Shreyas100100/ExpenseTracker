package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shreyas100100/ExpenseTracker/internal/routes"
	"github.com/shreyas100100/ExpenseTracker/pkg/database"
)

func main() {
	database.InitDB()
	log.Println("Database Initialized successfully")

	router := routes.SetupRouter()

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Printf("Server startting on port %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
