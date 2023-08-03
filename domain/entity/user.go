package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        string `gorm:"type:varchar;primary_key;"`
	Firstname string `json:"firstname" gorm:"not null;`
	Lastname  string `json:"lastname" gorm:"not null;`
	Email string `json:"email" gorm:"unique_index"`
	Tasks     Tasks
}

type Users []User
