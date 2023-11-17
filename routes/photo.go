package routes

import (
	"github.com/gorilla/mux"
	"github.com/drunkenbtterfly/task-5-pbi-btpns/controllers/photocontroller"
	"github.com/drunkenbtterfly/task-5-pbi-btpns/middleware"
)

func PhotoRoutes(router *mux.Router) {
	photo := router.PathPrefix("/photos").Subrouter()

	// Middleware
	photo.Use(middleware.Auth)

	photo.HandleFunc("", photocontroller.ListPhoto).Methods("GET")
	photo.HandleFunc("", photocontroller.CreatePhoto).Methods("POST")
	photo.HandleFunc("/{photoId}", photocontroller.ShowDetailPhoto).Methods("GET")
	photo.HandleFunc("/{photoId}", photocontroller.UpdatePhoto).Methods("PUT")
	photo.HandleFunc("/{photoId}", photocontroller.DeletePhoto).Methods("DELETE")
}
