package service

import (
	"assignment3/adapters"
	"assignment3/models"
	"assignment3/repositories"
	"strconv"
	"time"
)

type TaskService interface {
	CreateTask(task *models.Task) error
	GetTask(id int) (*models.Task, error)
	ClearCache() error
}

type taskService struct {
	repo   repositories.TaskRepository
	redis  repositories.RedisRepository
	logger adapters.ILogger
}

func NewService(repo repositories.TaskRepository, redis repositories.RedisRepository, logger adapters.ILogger) TaskService {
	return taskService{
		repo:   repo,
		redis:  redis,
		logger: logger,
	}
}

func (s taskService) CreateTask(task *models.Task) error {
	return s.repo.CreateTask(task)
}

func (s taskService) GetTask(id int) (*models.Task, error) {
	rvalue, err := s.redis.Get(strconv.Itoa(id), &models.Task{})
	if err == nil {
		s.logger.Info("Getting from cache", rvalue)
		return rvalue.(*models.Task), nil
	}
	task, err := s.repo.GetTask(id)
	if err != nil {
		s.logger.Error("Error getting task", err)
		return nil, err
	}

	s.logger.Info("Getting from database", task)

	err = s.redis.Set(strconv.Itoa(task.ID), task, time.Hour*2)
	if err != nil {
		s.logger.Error("Error setting cache", err)
	}

	return task, nil
}

func (s taskService) ClearCache() error {
	return s.redis.Clear()
}
