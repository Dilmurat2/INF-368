package repository

import (
	"assignment2/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "citizix_user"
	password = "S3cret"
	dbname   = "citizix_db"
)

type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetTask(id uint) (*models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(task *models.Task) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository() (TaskRepository, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected!")
	db.AutoMigrate(&models.Task{})
	return &repository{db: db}, nil
}

func (r *repository) CreateTask(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *repository) GetTask(id uint) (*models.Task, error) {
	task := &models.Task{}
	err := r.db.First(task, id).Error
	return task, err
}

func (r *repository) UpdateTask(task *models.Task) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(task).Update("Completed", true).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *repository) DeleteTask(task *models.Task) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Delete(task).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
