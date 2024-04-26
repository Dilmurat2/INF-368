package handlers

import (
	"assignment3/adapters"
	"assignment3/service"
)

type Handler struct {
	service service.TaskService
	logger  adapters.ILogger
}

func NewHandler(service service.TaskService, logger adapters.ILogger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}
