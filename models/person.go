package models

//import "github.com/jinzhu/gorm"

type Person struct {
	//gorm.Model
    
	ID       int64  `json:"id" gorm:"primary_key;auto_increment"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}
