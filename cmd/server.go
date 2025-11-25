package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/kodekage/gamma_mobility/internal/logger"
	"github.com/kodekage/gamma_mobility/utils"
)

func StartServer() {
	utils.EnvironmentSetup()
	logger.Info("Starting Gamma Mobility Server...")

	// mount routes
	router := setupRoutes()

	// start HTTP server
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_ADDRESS"), router))
}
