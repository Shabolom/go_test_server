package repository

import (
	"awesomeProject/config"
	"awesomeProject/iternal/domain"
)

type TeachersRepo struct {
}

func NewTeachersRepo() *TeachersRepo {
	return &TeachersRepo{}
}

func (tr *TeachersRepo) Save(teacher domain.Teachers) (domain.Teachers, error) {

	err := config.DB.Create(&teacher).Error

	if err != nil {
		return domain.Teachers{}, err
	}
	return teacher, nil
}

func (tr *TeachersRepo) Get() ([]domain.Teachers, error) {

	var teacher []domain.Teachers

	err := config.DB.Find(&teacher).Error

	if err != nil {
		return []domain.Teachers{}, err
	}
	return teacher, nil
}

func (tr *TeachersRepo) GetByID(key string) (domain.Teachers, error) {

	var teacher domain.Teachers

	err := config.DB.Where("id = ?", key).First(&teacher).Error

	if err != nil {
		return domain.Teachers{}, err
	}

	return teacher, nil
}
