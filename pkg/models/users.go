package models

import (
	"go/pkg/config"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt time.Time
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}
func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u

}
func GetUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users

}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("id=?", Id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(Id int64) User {
	var user User
	db.Where("id=?", Id).Delete(user)
	return user

}
