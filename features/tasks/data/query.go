package data

import "gorm.io/gorm"

type TaskQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TaskQuery {
	return &TaskQuery{
		db: db,
	}
}
