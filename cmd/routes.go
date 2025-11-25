package cmd

import (
	"github.com/gorilla/mux"
	paymentcontroller "github.com/kodekage/gamma_mobility/api/controllers"
)

func setupRoutes() *mux.Router {
	router := mux.NewRouter()

	{
		paymentcontroller.Mount(router)
	}

	return router
}
