package service

import (
	"github.com/ValentinAltamirano1/WashUp-Api/model"

	"errors"
)

type UserParams struct {
	Username     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(ur model.UserClient, userParams UserParams) (*model.User, error) {
	collisionUser, err := ur.UserFirst("email = ?", userParams.Email)
	if collisionUser != nil {
		return nil, errors.New("user already exists")
	}

	user := &model.User{
		Name:     userParams.Username,
		Email:    userParams.Email,
		Password: hashPassword(userParams.Password),
	}

	err = ur.SaveUser(user)
	if err != nil {
		return nil, errors.New("error trying to save user")
	}

	return user, nil
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(ur model.UserClient, loginParams LoginParams) (*LoginParams, error) {
	user, err := ur.UserFirst("email = ?", loginParams.Email)
	if err != nil {
		return nil, errors.New("error trying to find user")
	}

	err = ValidatePassword(user.Password, loginParams.Password)
	if err != nil {
		return nil, errors.New("invalid password")
	}
	
	return &LoginParams{
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
