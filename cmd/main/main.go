package main

import (
	"log"
	"net/http"

	"pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegistrarUsuarioRoute(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))

}
