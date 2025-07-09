package repository

import "SI/model"

func (r *Repository) FindAllUser() ([]model.User, error) {
	var users []model.User
	if err := r.DB.Model(&model.User{}).Find(&users).Error; err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func (r *Repository) FindOneUser(id string) (model.User, error) {
	var user model.User
	if err := r.DB.Where("id = ?", id).Find(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *Repository) FindUserByEmail(email string) (model.User, error) {
	var user model.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *Repository) CreateUser(user model.User) (model.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *Repository) UpdateUser(user model.User) (model.User, error) {
	if err := r.DB.Where("id = ?", user.Id).Updates(user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *Repository) DeleteUser(id string) error {
	var user model.User
	if err := r.DB.Where("id = ?", id).Delete(user).Error; err != nil {
		return err
	}

	return nil
}
