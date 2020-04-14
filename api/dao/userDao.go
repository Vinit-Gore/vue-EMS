package dao

import (
	"EventManagement/api/helpers"
	"EventManagement/api/models"

	"gopkg.in/mgo.v2/bson"
)

// GetAllUsersFromDB - Mongo
func GetAllUsersFromDB() error {
	res := []models.User{}
	if err := helpers.Collection(0).Find(nil).All(&res); err != nil {
		return err
	}
	models.UserList = res
	return nil
}

// SaveUserToDB - Mongo
func SaveUserToDB(user models.User) error {
	return helpers.Collection(0).Insert(user)
}

// DeleteUserFromDB - Mongo
func DeleteUserFromDB(id string) error {
	return helpers.Collection(0).Remove(bson.M{"loginID": id})
}

// UpdateUserFromDB - Mongo
// func UpdateUserFromDB(user models.User) error {
// user:= models.User{}
// 	return helpers.Collection(0).Update(bson.M{"$set": bson.M{"task": user.Task}})
// }
