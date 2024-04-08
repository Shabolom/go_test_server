package service

// слой работы с обработанными данными
import (
	"awesomeProject/iternal/domain"
	"awesomeProject/iternal/models"
	"awesomeProject/iternal/repository"
	"awesomeProject/iternal/tools"
	"github.com/gofrs/uuid"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

var userRepo = repository.NewUserRepo()

func (us UserService) Save(modelUser models.SaveUser) (domain.Users, error) {

	id, _ := uuid.NewV4()
	password, _ := tools.HashPassword(modelUser.Password)

	userEntity := domain.Users{
		Login:    modelUser.Login,
		Password: password,
	}

	userEntity.ID = id

	result, err := userRepo.Save(userEntity)

	if err != nil {
		return domain.Users{}, err
	}

	return result, nil
}

func (us UserService) Get() ([]domain.Users, error) {

	result, err := userRepo.Get()

	if err != nil {
		return []domain.Users{}, err
	}

	return result, nil
}

func (us UserService) GetByKey(key string) (domain.Users, error) {

	result, err := userRepo.GetByKey("id", key)

	if err != nil {
		return domain.Users{}, err
	}

	return result, nil
}

func (us UserService) Update(key string, body models.SaveUser) (domain.Users, error) {

	result, err := userRepo.Update(key, body)

	if err != nil {
		return domain.Users{}, err
	}

	return result, nil
}

func (us UserService) Delet(key, value string) (domain.Users, error) {

	result, err := userRepo.Delet(key, value)

	if err != nil {
		return domain.Users{}, err
	}

	return result, nil
}
