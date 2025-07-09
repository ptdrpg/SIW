package controller

import (
	"SI/lib"
	"SI/model"

	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutput struct {
	Data model.LoginResponse `json:"data"`
}

func (c *Controller) Login(input LoginInput) (*LoginOutput, error) {
	user, err := c.R.FindUserByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, err
	}

	token, err := lib.GenerateToken(user.Name)
	if err != nil {
		return nil, err
	}

	login := model.LoginResponse{
		User:  user,
		Token: token,
	}
	res := LoginOutput{
		Data: login,
	}

	return &res, nil
}
