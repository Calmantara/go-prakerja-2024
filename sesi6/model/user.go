package model

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

var Users = []*User{}

type User struct {
	ID       uint64 `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
	Email    string `json:"email" gorm:"column:email"`
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("invalid username input")
	}

	if u.Email == "" {
		return errors.New("invalid username input")
	}
	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	log.Println("THIS IS FROM BEFORE CREATE HOOK")
	return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	log.Println("THIS IS FROM AFTER CREATE HOOK")
	return
}
