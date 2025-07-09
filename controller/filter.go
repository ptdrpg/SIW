package controller

import (
	"SI/model"

)

func (c *Controller) FilterConfirmation() (*Members, error) {

	allMembers, err := c.R.FindAllMember()
	if err != nil {
		return nil, err
	}

	var confirm []model.Member

	for i := 0; i < len(allMembers); i++ {
		if allMembers[i].Confirmation {
			confirm = append(confirm, allMembers[i])
		}
	}

	res := Members{
		Data: confirm,
	}

	return  &res, nil
}

func (c *Controller) FilterCommunion() (*Members, error) {

	allMembers, err := c.R.FindAllMember()
	if err != nil {
		return nil, err
	}

	var communion []model.Member
	for i := 0; i < len(allMembers); i++ {
		if allMembers[i].Communion {
			communion = append(communion, allMembers[i])
		}
	}

	res := Members{
		Data: communion,
	}

	return  &res, nil
}

func (c *Controller) FilterConfession() (*Members, error) {
	allMembers, err := c.R.FindAllMember()
	if err != nil {
		return nil, err
	}

	var confession []model.Member
	for i := 0; i < len(allMembers); i++ {
		if allMembers[i].Confession {
			confession = append(confession, allMembers[i])
		}
	}

	res := Members{
		Data: confession,
	}

	return  &res, nil
}

func (c *Controller) FilterBaptised() (*Members, error) {
	allMembers, err := c.R.FindAllMember()
	if err != nil {
		return nil, err
	}

	var baptised []model.Member
	for i := 0; i < len(allMembers); i++ {
		if allMembers[i].Baptism {
			baptised = append(baptised, allMembers[i])
		}
	}

	res := Members{
		Data: baptised,
	}

	return &res, nil
}

func (c *Controller) FilterRenewBaptem() (*Members, error) {
	allMembers, err := c.R.FindAllMember()
	if err != nil {
		return nil, err
	}

	var renew []model.Member
	for i := 0; i < len(allMembers); i++ {
		if allMembers[i].BaptismRenew {
			renew = append(renew, allMembers[i])
		}
	}

	res := Members{
		Data: renew,
	}

	return  &res, nil
}

func (c *Controller) FilterMaried() (*Members, error) {
	allMembers, err := c.R.FindAllMember()
	if err != nil {
		return nil, err
	}

	var maried []model.Member
	for i := 0; i < len(allMembers); i++ {
		if allMembers[i].Wedding {
			maried = append(maried, allMembers[i])
		}
	}

	res := Members{
		Data: maried,
	}

	return  &res, nil
}
