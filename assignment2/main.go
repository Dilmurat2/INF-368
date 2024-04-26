package main

import (
	"assignment2/models"
	repository2 "assignment2/repository"
	"fmt"
	"log"
)

func main() {
	repo, err := repository2.NewRepository()
	if err != nil {
		log.Fatalln(err)
	}

	task := models.Task{
		ID:        5,
		Name:      "Pupa",
		Completed: false,
	}

	err = repo.CreateTask(&task)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Task created")
	getTask, err := repo.GetTask(task.ID)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(getTask)
	task.Completed = true
	if err := repo.UpdateTask(&task); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Task updated")
	task1, err := repo.GetTask(task.ID)
	if err != nil {
		return
	}

	fmt.Println(task1)
}
