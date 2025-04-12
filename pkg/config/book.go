package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var  db * gorm.DB

func ConnectBook() {
	d, err := gorm.Open("mysql","henry:Phucprohz123@@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err) 
	}
	db = d

}

func GetBook() *gorm.DB{
	return db
}