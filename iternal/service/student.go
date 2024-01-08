package service

import (
	"awesomeProject/iternal/domain"
	"awesomeProject/iternal/models"
	"awesomeProject/iternal/repository"
	"github.com/gofrs/uuid"
)

type StudentService struct {
}

func NewStudentService() *StudentService {
	return &StudentService{}
}

var studentRepo = repository.NewStudentRepo()

func (sr *StudentService) Save(studentModel models.SaveStudent) (domain.Student, error) {

	// создаем uuid
	id, _ := uuid.NewV4()

	// тут происходит создание объекта
	studentEntity := domain.Student{
		Name:    studentModel.Name,
		Surname: studentModel.Surname,
		Gender:  studentModel.Gender,
		Age:     studentModel.Age,
		Country: studentModel.Country,
		Email:   studentModel.Email,
	}

	// обновляем id
	studentEntity.ID = id

	// вызываем метод сохранения в репозитории
	result, err := studentRepo.Save(studentEntity)

	if err != nil {
		return domain.Student{}, err
	}

	return result, nil
}

func (sr *StudentService) Get() ([]domain.Student, error) {

	result, err := studentRepo.Get()

	if err != nil {
		return []domain.Student{}, err
	}

	return result, nil
}

func (sr *StudentService) GetByID(key string) (domain.Student, error) {

	result, err := studentRepo.GetByID(key)

	if err != nil {
		return domain.Student{}, err
	}

	return result, nil
}
