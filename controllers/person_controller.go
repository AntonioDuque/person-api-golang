package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AntonioDuque/demo-crud-api-rest-go/commons"
	"github.com/AntonioDuque/demo-crud-api-rest-go/models"
	"github.com/gorilla/mux"
)

// find all
func GetAll(w http.ResponseWriter, r *http.Request) {
	persons := []models.Person{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&persons)

	json, _ := json.Marshal(persons)

	commons.SendResponse(w, http.StatusOK, json)
}

// find one
func Get(w http.ResponseWriter, r *http.Request){
	person := models.Person{}

	id:= mux.Vars(r)["id"]

	db:= commons.GetConnection()
	defer db.Close()

	db.Find(&person, id)

	if person.ID >0{
		json,_ := json.Marshal(person)
		commons.SendResponse(w,http.StatusOK,json)
	} else {
		commons.SendError(w,http.StatusNotFound)
	} 

}


// save 
func Save(w http.ResponseWriter, r *http.Request){
    person:= models.Person{}
	
	db:= commons.GetConnection()
	defer db.Close()

	err:= json.NewDecoder(r.Body).Decode(&person)

	if err!= nil{
		log.Fatal(err)
		commons.SendError(w,http.StatusBadRequest)
		return
	}
 
	err= db.Save(&person).Error
	
	if err!= nil{
		log.Fatal(err)
		commons.SendError(w,http.StatusInternalServerError)
		return
	}

	json,_:= json.Marshal(person)

	commons.SendResponse(w, http.StatusCreated,json)


}

// delete 
func Delete(w http.ResponseWriter, r *http.Request){
	person:= models.Person{}

	db:= commons.GetConnection()
	defer db.Close()

	id:= mux.Vars(r)["id"]

	db.Find(&person,id)

	if person.ID > 0{
		db.Delete(person)
		commons.SendResponse(w,http.StatusOK,[]byte(`{}`) )
	} else{
		commons.SendError(w,http.StatusNotFound)
	}

}
