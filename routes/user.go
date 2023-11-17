package routes

import (
	"github.com/gorilla/mux"
	"github.com/drunkenbtterfly/task-5-pbi-btpns/controllers/usercontroller"
	"github.com/drunkenbtterfly/task-5-pbi-btpns/middleware"
)

func UserRoutes(router *mux.Router) {
	user := router.PathPrefix("/users").Subrouter()

	// Middleware
	user.Use(middleware.Auth)

	user.HandleFunc("", usercontroller.ListUser).Methods("GET")
	user.HandleFunc("/{userId}", usercontroller.UpdateUser).Methods("PUT")
	user.HandleFunc("/{userId}", usercontroller.DeleteUser).Methods("DELETE")
}
