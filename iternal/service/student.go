package service

import (
	"awesomeProject/domain"
	"awesomeProject/iternal/models"
	"awesomeProject/iternal/repository"
	"fmt"
	"github.com/gofrs/uuid"
)

type StudentService struct {
}

func NewStudentService() *StudentService { return &StudentService{} }

var studentRepo = repository.NewStudentRepo()

func (sr *StudentService) Save(studentModel models.SaveStudent) (domain.Student, error) {

	id, _ := uuid.NewV4()

	studentEntity := domain.Student{
		Name:    studentModel.Name,
		Surname: studentModel.Surname,
		Gender:  studentModel.Gender,
		Age:     studentModel.Age,
		Country: studentModel.Country,
		Email:   studentModel.Email,
	}
	studentEntity.ID = id
	fmt.Println(studentEntity)
	result, err := studentRepo.Save(studentEntity)
	fmt.Println(studentEntity)
	if err != nil {
		return domain.Student{}, err
	}

	return result, nil
}
