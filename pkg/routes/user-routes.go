package routes

import (
	"go/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegistrarUsuarioRoute = func(router *mux.Router) {
	router.HandleFunc("/usuario/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/usuario/", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/usuario/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/usuario/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/usuario/{id}", controllers.DeleteUser).Methods("DELETE")
}
