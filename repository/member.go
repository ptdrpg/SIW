package repository

import "SI/model"

func (r *Repository) FindAllMember() ([]model.Member, error) {
	var members []model.Member
	if err := r.DB.Model(&model.Member{}).Find(&members).Error; err != nil {
		return []model.Member{}, err
	}

	return members, nil
}

func (r *Repository) FindOneMember(id string) (model.Member, error) {
	var member model.Member
	if err := r.DB.Where("id = ?", id).Find(&member).Error; err != nil {
		return model.Member{}, err
	}

	return member, nil
}

func (r *Repository) CreateMember(member model.Member) (model.Member, error) {
	if err := r.DB.Create(member).Error; err != nil {
		return model.Member{}, err
	}

	return member, nil
}

func (r *Repository) UpdateMember(member model.Member) (model.Member, error) {
	if err := r.DB.Where("id = ?", member.Id).Updates(member).Error; err != nil {
		return model.Member{}, err
	}

	return member, nil
}

func (r *Repository) DeleteMember(id string) error {
	var member model.Member
	if err := r.DB.Where("id = ?", id).Delete(&member).Error; err != nil {
		return err
	}

	return nil
}
