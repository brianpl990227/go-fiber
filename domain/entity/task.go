package entity

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Id          string `gorm:"type:varchar;primary_key;"`
	Title       string `gorm:"type:varchar(100);not null;unique_index"`
	Description string
	Done        bool `gorm:"default:false"`
	UserID      string
}

type Tasks []Task
