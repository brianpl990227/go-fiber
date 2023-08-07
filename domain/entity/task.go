package entity

import "gorm.io/gorm"


type Task struct {
	gorm.Model
	ID          string `json:"id" gorm:"primaryKey`
	Title       string `json:"title" gorm:"type:varchar(100);not null;unique_index`
	Description string `json:"description"`
	Done        bool `gorm:"default:false" json:"done"`
	UserID      string `json:"userId"`
}


