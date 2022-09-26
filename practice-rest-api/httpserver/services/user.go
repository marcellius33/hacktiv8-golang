package services

import (
	"golang.org/x/crypto/bcrypt"
	"practice-rest-api/httpserver/controllers/params"
	"practice-rest-api/httpserver/controllers/views"
	"practice-rest-api/httpserver/repositories/models"
)

func CreateUser(req *params.UserCreateRequest) *views.Response {
	var user models.User

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return views.BadRequestError(err)
	}

	user.ID = 1
	user.Password = string(hash)
	user.Username = req.Username

	v := views.SuccessCreateResponse(user, "created success!")

	return v
}
