package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kodekage/gamma_mobility/internal/logger"
)

func StartServer() {
	environmentSetup()
	logger.Info("Starting Gamma Mobility Server...")

	// setup routes
	router := setupRoutes()

	// start HTTP server
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_ADDRESS"), router))
}

func environmentSetup() {
	if err := godotenv.Load(); err != nil {
		logger.Warn("No .env file found, proceeding with system environment variables")
	}

	logger.Info("Environment variables loaded")
}
