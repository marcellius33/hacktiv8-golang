package services

import (
	"clean-architecture-1/httpserver/controllers/params"
	"clean-architecture-1/httpserver/controllers/views"
	"clean-architecture-1/httpserver/repositories/models"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
)

var users = []models.User{
	{
		ID:       1,
		Email:    "marcel@gmail.com",
		Password: "secret123",
	},
	{
		ID:       2,
		Email:    "curry@gmail.com",
		Password: "secret123@",
	},
	{
		ID:       3,
		Email:    "stephen@gmail.com",
		Password: "splash123",
	},
}

func CreateUser(req *params.UserCreateRequest) *views.Response {
	var user models.User

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return views.BadRequestError(err)
	}

	user.ID = rand.Int()
	user.Password = string(hash)
	user.Email = req.Email

	users = append(users, user)

	v := views.SuccessCreateResponse(user, "created success!")

	return v
}

func GetUsers() *views.Response {
	v := views.SuccessCreateResponse(users, "get user success!")
	return v
}
