package router

import (
	"golangblogbackend/controllers"
	"golangblogbackend/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/signup", controllers.Signup).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", controllers.Login).Methods("POST", "OPTIONS")

	authenticatedRouter := router.PathPrefix("/api").Subrouter()
	authenticatedRouter.Use(middleware.Authenticate)
	authenticatedRouter.HandleFunc("/posts", controllers.GetAllPosts).Methods("GET", "OPTIONS")
	authenticatedRouter.HandleFunc("/posts/{id}", controllers.GetPostByID).Methods("GET", "OPTIONS")
	authenticatedRouter.HandleFunc("/posts", controllers.CreatePost).Methods("POST", "OPTIONS")
	authenticatedRouter.HandleFunc("/posts/{id}", controllers.UpdatePost).Methods("PUT", "OPTIONS")
	authenticatedRouter.HandleFunc("/posts/{id}", controllers.DeletePost).Methods("DELETE", "OPTIONS")

	return router
}
