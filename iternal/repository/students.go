package repository

import (
	"awesomeProject/config"
	"awesomeProject/domain"
)

type StudentRepo struct {
}

func NewStudentRepo() *StudentRepo {
	return &StudentRepo{}
}

func (sr *StudentRepo) Save(student domain.Student) (domain.Student, error) {

	err := config.DB.Create(student).Error

	if err != nil {
		return domain.Student{}, err
	}
	return student, nil
}
