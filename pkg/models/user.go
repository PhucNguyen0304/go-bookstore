
package models

import (
	"github.com/jinzhu/gorm"
)
type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email`
	Password string `json:"pass`
}



func (u *User) RegisterUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetAllUser() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(ID int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("id = ?", ID).Find(&getUser)
	return &getUser,db
}

func DeleteUser(ID int64) *User {
	var User User
	db.Where("id = ?", ID).Delete(User)
	return &User
}

func GetUserByEmail(email string) (*User, error) {
    var user User
    err := db.Where("email = ?", email).First(&user).Error
    return &user, err 
}