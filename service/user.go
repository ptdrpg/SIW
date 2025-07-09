package service

import (
	"SI/lib"
	"SI/model"

	"github.com/google/uuid"
)

func (s *Service) CreateUserAccount(input model.UserInput) (model.LoginResponse, error) {
	user := model.User{
		Id:        uuid.New().String(),
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}

	res, err := s.R.CreateUser(user)
	if err != nil {
		return model.LoginResponse{}, err
	}

	token, err := lib.GenerateToken(res.Name)
	if err != nil {
		return model.LoginResponse{}, err
	}
	response := model.LoginResponse{
		User:  res,
		Token: token,
	}
	return response, nil
}
