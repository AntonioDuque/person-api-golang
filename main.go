package main

import (
	"log"
	"net/http"

	"github.com/AntonioDuque/demo-crud-api-rest-go/commons"
	"github.com/AntonioDuque/demo-crud-api-rest-go/routes"

	"github.com/gorilla/mux"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()
    routes.SetPersonRoutes(router)
 
	server:= http.Server{
		Addr: ":9000",
		Handler: router,
	}

	log.Println("Server running --> port:9000")

	log.Println(server.ListenAndServe())

}