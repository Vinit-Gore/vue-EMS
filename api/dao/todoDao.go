package dao

import (
	"EventManagement/api/helpers"
	"EventManagement/api/models"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// GetAllTodosFromDB - Mongo
func GetAllTodosFromDB() error {
	res := []models.Todo{}
	if err := helpers.Collection(1).Find(nil).All(&res); err != nil {
		return err
	}
	models.TodoList = res
	return nil
}

// SaveTodoToDB - Mongo
func SaveTodoToDB(todo models.Todo) error {
	uuidVar, _ := uuid.NewV4() // 47891-sskjsknfdsf-89624-dhnkskfhdkf
	todo.Id = uuidVar.String()
	return helpers.Collection(1).Insert(todo)
}

// DeleteTodoFromDB - Mongo
func DeleteTodoFromDB(id string) error {
	return helpers.Collection(1).Remove(bson.M{"_id": id})
}

// UpdateTodoFromDB - Mongo
func UpdateTodoFromDB(todo models.Todo) error {
	return helpers.Collection(1).Update(bson.M{"_id": todo.Id}, bson.M{"$set": bson.M{"task": todo.Task}})
}
