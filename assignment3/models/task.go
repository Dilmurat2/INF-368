package models

type Task struct {
	ID        int    `json:"ID"`
	Name      string `json:"Name"`
	Completed bool   `json:"Completed"`
}
