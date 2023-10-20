package service

import (
	"errors"

	"github.com/ValentinAltamirano1/WashUp-Api/model"
	"golang.org/x/crypto/bcrypt"
)

type UserParams struct {
	Username     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(ur model.UserClient, userParams UserParams) (*model.User, error) {
	collisionUser, err := ur.UserFirst("email = ?", userParams.Email)
	if collisionUser != nil {
		return nil, errors.New("email already exists")
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

func hashPassword(plainPassword string) string {
	bcryptHash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 12)
	if err != nil {
		return "error"
	}
	return string(bcryptHash)
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

func ValidatePassword(passwordHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}
