package router

import (
	"github.com/divy-11/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.AllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreatingMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/movie/{id}", controller.WatchedMovie).Methods("PUT")
	router.HandleFunc("/api/movies", controller.AllMovies).Methods("DELETE")
	return router
}
