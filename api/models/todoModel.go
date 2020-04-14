package models

type Todo struct {
	Id   string    `json:"id"`
	Task string `json:"task"`
}

// Created global list which will be available throughout the application
var TodoList []Todo
