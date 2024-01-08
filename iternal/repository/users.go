package repository

import (
	"awesomeProject/config"
	"awesomeProject/iternal/domain"
	"awesomeProject/iternal/models"
)

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) Save(user domain.Users) (domain.Users, error) {

	err := config.DB.
		Create(&user).Error

	if err != nil {
		return domain.Users{}, err
	}
	return user, nil
}

func (ur *UserRepo) Get() ([]domain.Users, error) {

	var user []domain.Users

	err := config.DB.
		Find(&user).Error

	if err != nil {
		return []domain.Users{}, err
	}
	return user, nil
}

func (ur *UserRepo) Update(key string, body models.SaveUser) (domain.Users, error) {

	var user domain.Users

	err := config.DB.
		Where("id = ?", key).
		Updates(domain.Users{
			Login:    body.Login,
			Password: body.Password,
		}).
		First(&user).
		Error

	if err != nil {
		return domain.Users{}, err
	}

	return user, nil
}

// GetByKey возвращает пользователя по ключу
func (ur *UserRepo) GetByKey(key, value string) (domain.Users, error) {
	var user domain.Users
	err := config.DB.
		Unscoped().
		Where(key+" = ?", value).
		First(&user).Error

	return user, err
}
