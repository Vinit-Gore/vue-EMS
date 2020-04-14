package services

import (
	"EventManagement/api/dao"
	"EventManagement/api/models"
)

func NewTodo() models.Todo {
	return models.Todo{}
}

func GetTodoList() []models.Todo {
	if models.TodoList == nil || len(models.TodoList) == 0 {
		models.TodoList = []models.Todo{}
		return models.TodoList
	}
	return models.TodoList
}

// AddTodoService Add Todo Service
// func AddTodoService(todo models.Todo) {
// 	tList := GetTodoList()
// 	tList = append(tList, todo)
// 	models.TodoList = tList
// }

// // UpdateTodoService Update Todo Service
// func UpdateTodoService(todo models.Todo) {
// 	tList := GetTodoList()
// 	for i := 0; i <= len(tList); i++ {
// 		if tList[i].Id == todo.Id {
// 			tList[i].Task = todo.Task
// 			break
// 		}
// 	}
// 	models.TodoList = tList
// }

// // DeleteTodoService Update Todo Service
// func DeleteTodoService(todo models.Todo) {
// 	tList := GetTodoList()
// 	for i := 0; i <= len(tList); i++ {
// 		if tList[i].Id == todo.Id {
// 			tList = append(tList[:i], tList[i+1:]...)
// 			break
// 		}
// 	}
// 	models.TodoList = tList
// }

func AddTodoService(todo models.Todo) {
	// Mongo DB
	dao.SaveTodoToDB(todo)
	dao.GetAllTodosFromDB()

	//MYSQL
	// SaveTodoToMYSQLDB(todo)
	// GetAllTodosFromMYSQLDB()
}

// UpdateTodoService Update Todo Service
func UpdateTodoService(todo models.Todo) {
	// MongoDB
	dao.UpdateTodoFromDB(todo)
	dao.GetAllTodosFromDB()

	//MySQL
	// UpdateTodoFromMYSQLDB(todo)
	// GetAllTodosFromMYSQLDB()

}

// DeleteTodoService Update Todo Service
func DeleteTodoService(todo models.Todo) {
	// MongoDB
	dao.DeleteTodoFromDB(todo.Id)
	dao.GetAllTodosFromDB()

	//MYSQL
	// DeleteTodoFromMYSQLDB(todo.Id)
	// GetAllTodosFromMYSQLDB()

}
