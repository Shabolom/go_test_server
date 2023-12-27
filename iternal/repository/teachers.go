package repository

import (
	"awesomeProject/config"
	"awesomeProject/domain"
	"github.com/gofrs/uuid"
)

type TeachersRepo struct {
}

func NewTeachersRepo() *TeachersRepo {
	return &TeachersRepo{}
}

func (sr *TeachersRepo) Save() (domain.Teachers, error) {

	user := domain.Teachers{
		Base:      domain.Base{},
		Name:      "Dana",
		Surname:   "Dancov",
		MidleName: "Sharapov",
	}
	user.ID, _ = uuid.NewV4()
	err := config.DB.Create(&user).Error

	if err != nil {
		return domain.Teachers{}, err
	}
	return user, nil
}
