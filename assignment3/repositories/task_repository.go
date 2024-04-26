package repositories

import (
	"assignment3/config"
	"assignment3/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetTask(id int) (*models.Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(config *config.Config) (TaskRepository, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DB.Host, config.DB.Port, config.DB.Username, config.DB.Password, config.DB.Name)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&models.Task{}); err != nil {
		return nil, err
	}
	return &taskRepository{db: db}, nil
}

func (r *taskRepository) CreateTask(task *models.Task) error {
	return r.db.Create(&task).Error
}

func (r *taskRepository) GetTask(id int) (*models.Task, error) {
	var result models.Task
	query := "SELECT * FROM tasks WHERE id = ?"
	if err := r.db.Raw(query, id).Scan(&result).Error; err != nil {
		return nil, err
	}
	if result.ID == 0 {
		return nil, fmt.Errorf("task not found")
	}
	return &result, nil
}
