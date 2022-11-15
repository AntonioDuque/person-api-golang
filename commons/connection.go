package commons

import (
	"log"

	"github.com/AntonioDuque/demo-crud-api-rest-go/models"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql","root:@/crud_gorm_golang_app")

	if err != nil {
		log.Fatal(err)
	}
	return db
}


func Migrate(){
	db:= GetConnection()
	defer db.Close()

	log.Println("Migrating")

	db.AutoMigrate(&models.Person{})
}