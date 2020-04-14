package models

type User struct {
	Name     string   `json:"name"`
	LoginID  string   `json:"loginID"`
	Password string   `json:"password"`
	Group    []string `json:"group"`
}

var UserList []User

type Login struct {
	LoginID  string `json:"loginID"`
	Password string `json:"password"`
}
