package main

import (
	"assignment3/adapters"
	"assignment3/config"
	"assignment3/handlers"
	"assignment3/repositories"
	"assignment3/service"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	logger := adapters.NewLogger()
	rds, err := repositories.NewRedisClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	taskRepo, err := repositories.NewTaskRepository(cfg)
	if err != nil {
		log.Fatal(err)
	}

	taskService := service.NewService(taskRepo, rds, logger)
	handler := handlers.NewHandler(taskService, logger)

	app := gin.Default()

	app.POST("/task", handler.CreateTask)
	app.GET("/task/:id", handler.GetTask)
	app.DELETE("/clear-cache", handler.ClearCache)

	app.Run(":8080")
}
