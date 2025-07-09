package controller

import (
	"SI/model"

	"github.com/google/uuid"
)

type Members struct {
	Data []model.Member `json:"data"`
}

type Member struct {
	Data model.Member `json:"data"`
}

func (c *Controller) FindAllMember() (*Members, error) {
	members, err := c.R.FindAllMember()
	if err != nil {
		return nil, err
	}

	res := Members{
		Data: members,
	}

	return &res, nil
}

func (c *Controller) FindOneMember(id string) (*Member, error) {

	member, err := c.R.FindOneMember(id)
	if err != nil {
		return nil, err
	}

	res := Member{
		Data: member,
	}

	return &res, nil
}

func (c *Controller) CreateMember(input model.Member) (*Member, error) {
	if !input.Wedding {
		input.Couple = " "
	}

	input.Id = uuid.New().String()
	member, err := c.R.CreateMember(input)
	if err != nil {
		return nil, err
	}

	res := Member{
		Data: member,
	}
	return &res, nil
}

func (c *Controller) UpdateMember(id string, input model.Member) (*Member, error) {
	input.Id = id
	member, err := c.R.UpdateMember(input)
	if err != nil {
		return nil, err
	}

	res := Member{
		Data: member,
	}

	return &res, nil
}

func (c *Controller) DeleteMember(id string) (*Deleteresponse, error){
	err := c.R.DeleteMember(id)
	if err != nil {
		return nil, err
	}

	response := Deleteresponse{
		Data: "member succefuly deleted",
	}

	return  &response, nil
}
