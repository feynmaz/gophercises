package storage

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Description string
	Done        bool
}

func DoMigrations() {
	db.AutoMigrate(&Task{})
}
