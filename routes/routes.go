package routes

import (
	"net/http"
	"tengrific/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id:[0-9]+}", controllers.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/users", controllers.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id:[0-9]+}", controllers.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id:[0-9]+}", controllers.DeleteUser).Methods(http.MethodDelete)

	return router
}
