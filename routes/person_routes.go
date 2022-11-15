package routes

import (
	"github.com/AntonioDuque/demo-crud-api-rest-go/controllers"
	"github.com/gorilla/mux"
)

func SetPersonRoutes(router *mux.Router){
	subRoute:= router.PathPrefix("/person/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/save",controllers.Save).Methods("POST")
	subRoute.HandleFunc("/delete/{id}",controllers.Delete).Methods("DELETE")
	subRoute.HandleFunc("/find/{id}",controllers.Get).Methods("GET")
}