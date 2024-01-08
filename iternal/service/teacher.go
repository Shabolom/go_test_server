package service

import (
	"awesomeProject/iternal/domain"
	"awesomeProject/iternal/models"
	"awesomeProject/iternal/repository"
	"github.com/gofrs/uuid"
)

type TeacherService struct {
}

func NewTeacherService() *TeacherService {
	return &TeacherService{}
}

var teacherRepo = repository.NewTeachersRepo()

func (ts TeacherService) Save(teacherModel models.SaveTeacher) (domain.Teachers, error) {

	id, _ := uuid.NewV4()

	teacherEntity := domain.Teachers{
		Name:      teacherModel.Name,
		Surname:   teacherModel.Surname,
		MidleName: teacherModel.MidleName,
	}
	teacherEntity.ID = id

	result, err := teacherRepo.Save(teacherEntity)

	if err != nil {
		return domain.Teachers{}, err
	}
	return result, nil
}

func (ts TeacherService) Get() ([]domain.Teachers, error) {

	result, err := teacherRepo.Get()

	if err != nil {
		return []domain.Teachers{}, err
	}

	return result, nil
}

func (ts TeacherService) GetByID(key string) (domain.Teachers, error) {

	result, err := teacherRepo.GetByID(key)

	if err != nil {
		return domain.Teachers{}, err
	}

	return result, nil
}
