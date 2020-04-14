package services

import (
	"EventManagement/api/dao"
	"EventManagement/api/models"
)

func GetUserList() []models.User {
	if models.UserList == nil || len(models.UserList) == 0 {
		models.UserList = []models.User{}
		return models.UserList
	}
	return models.UserList
}

func GetUser(user models.User) models.User {
	for index := 0; index < len(models.UserList); index++ {
		if models.UserList[index].LoginID == user.LoginID && models.UserList[index].Password == user.Password {
			return models.UserList[index]
		}
	}
	return models.User{}
}

func AddUsers(user models.User) []models.User {
	models.UserList = append(models.UserList, user)
	return models.UserList
}

func ValidatedUser(userRequest models.User) (models.User, bool, error) {
	user := GetUser(userRequest)
	if user.LoginID != "" {
		return user, true, nil
	}
	return models.User{}, false, nil
}

func AddUserService(todo models.User) {
	// Mongo DB
	dao.SaveUserToDB(todo)
	dao.GetAllUsersFromDB()

	//MYSQL
	// SaveUserToMYSQLDB(todo)
	// GetAllUsersFromMYSQLDB()
}

// UpdateUserService Update User Service
func UpdateUserService(todo models.User) {
	// MongoDB
	dao.UpdateUserFromDB(todo)
	dao.GetAllUsersFromDB()

	//MySQL
	// UpdateUserFromMYSQLDB(todo)
	// GetAllUsersFromMYSQLDB()

}

// DeleteUserService Update User Service
func DeleteUserService(todo models.User) {
	// MongoDB
	dao.DeleteUserFromDB(todo.Id)
	dao.GetAllUsersFromDB()

	//MYSQL
	// DeleteUserFromMYSQLDB(todo.Id)
	// GetAllUsersFromMYSQLDB()

}
