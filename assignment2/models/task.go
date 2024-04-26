package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID        uint
	Name      string
	Completed bool
}
