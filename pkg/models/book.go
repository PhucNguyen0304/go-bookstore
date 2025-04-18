package models

import (
	"github.com/jinzhu/gorm"
	"github.com/PhucNguyen0304/go-bookstore/pkg/config"

)
var db *gorm.DB
var jwtSecret []byte


type Book struct {
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author`
	Publication string `json:"publication`
}

func init() {
	config.ConnectBook()
	db = config.GetBook()
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Cart{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, db
} 

func DeleteBook(Id int64) *Book {
	var book Book
	 db.Where("ID=?", Id).Delete(book)
	return &book
}