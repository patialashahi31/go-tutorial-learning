package router

import (
	"mongoapi/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAPIAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateOneMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/delete", controller.DeleteAllMovies).Methods("DELETE")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")

	return router
}