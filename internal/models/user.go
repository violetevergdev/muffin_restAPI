package models

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var UserMock = User{
	Username: "1",
	Password: "1",
}