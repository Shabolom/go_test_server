package repository

import (
	"awesomeProject/config"
	"awesomeProject/iternal/domain"
)

type StudentRepo struct {
}

func NewStudentRepo() *StudentRepo {
	return &StudentRepo{}
}

// Save принимает объект структуры students (те уже заполненный объект на основе структуры)
// по сути мы тут делаем insert из SQL
func (sr *StudentRepo) Save(student domain.Student) (domain.Student, error) {

	// метод библиотеки для сохранения сущности в базе данных
	err := config.DB.Create(&student).Error

	if err != nil {
		return domain.Student{}, err
	}

	return student, nil
}

func (sr *StudentRepo) Get() ([]domain.Student, error) {

	var student []domain.Student

	err := config.DB.
		Find(&student).
		Error

	if err != nil {
		return []domain.Student{}, err
	}

	return student, nil
}

func (sr *StudentRepo) GetByID(key string) (domain.Student, error) {

	var student domain.Student

	err := config.DB.
		Where("id = ?", key).
		First(&student).
		Error

	if err != nil {
		return domain.Student{}, err
	}

	return student, nil

	//var url domain.URL
	//if err := config.DB.
	//	Table("urls as u").
	//	Select("u.*").
	//	Where("u.short_url = ?", shortURL).
	//	Scan(&url).
	//	Error; err != nil {
	//	return domain.URL{}, err
	//}
	//return url, nil
}
