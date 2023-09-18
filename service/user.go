package service

import (
	"errors"

	"github.com/ValentinAltamirano1/WashUp-Api/model"
)

type UserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(ur model.UserClient, userParams UserParams) error {
	collisionUser, err := ur.UserFirst("email = ?", userParams.Email)
	if collisionUser != nil {
		return errors.New("email already exists")
	}

	user := &model.User{
		Name:     userParams.Name,
		Email:    userParams.Email,
		Password: userParams.Password,
	}

	err = ur.SaveUser(user)
	if err != nil {
		return errors.New("error trying to save user")
	}

	return nil
}
