package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        string `gorm:"primary_key;" json:"id"`
	Firstname string `json:"firstname" gorm:"not null;`
	Lastname  string `json:"lastname" gorm:"not null;`
	Email     string `json:"email" gorm:"unique_index"`
	Tasks     []Task 
}

