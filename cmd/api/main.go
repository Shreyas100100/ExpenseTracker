package main

import (
	"net/http"

	"github.com/shreyas100100/ExpenseTracker/internal/routes"
	"github.com/shreyas100100/ExpenseTracker/pkg/database"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	database.InitDB()
	logrus.Info("Database initialized successfully")

	router := routes.SetupRouter()

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	logrus.Infof("Server starting on port %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Fatal("Server encountered an error")
	}
}
