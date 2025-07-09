package controller

import (
	"SI/model"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserList struct {
	Data []model.User `json:"data"`
}

type UserRes struct {
	Data model.User `json:"data"`
}

type NewUser struct {
	Data model.LoginResponse `json:"data"`
}

type Deleteresponse struct {
	Data string `json:"data"`
}

func (c *Controller) FindAllUser() (*UserList, error) {
	users, err := c.R.FindAllUser()
	if err != nil {
		return nil, err
	}

	response := UserList{
		Data: users,
	}

	return  &response, nil
}

func (c *Controller) FindOneUser(id string) (*UserRes, error) {
	user, err := c.R.FindOneUser(id)
	if err != nil {
		return nil, err
	}

	response := UserRes{
		Data: user,
	}

 return  &response, nil
}

func (c *Controller) CreateUser(input model.UserInput) (*NewUser, error) {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return nil, err
	}

	input.CreatedAt = time.Now().String()
	input.UpdatedAt = time.Now().String()
	input.Password = string(hashedPass)
	newUser, err := c.S.CreateUserAccount(input)
	if err != nil {
		return nil, err
	}
	res := NewUser{
		Data: newUser,
	}
	
	return  &res, nil
}

func (c *Controller) UpdateUser(id string, input model.UpdateUserInput) (*UserRes, error) {
	user, err := c.R.FindOneUser(id)
	if err != nil {
		return nil, err
	}

	user.UpdatedAt = time.Now().String()
	user.Email = input.Email
	user.Name = input.Name

	updatedName, err := c.R.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	response := UserRes{
		Data: updatedName,
	}

	return  &response, nil
}

func (c *Controller) DeleteUser(id string) (*Deleteresponse, error) {
	err := c.R.DeleteUser(id)
	if err != nil {
		return nil, err
	}

	response := Deleteresponse{
		Data: "user succefuly deleted",
	}

	return  &response, nil
}
