package api

import (
	"awesomeProject/iternal/models"
	"awesomeProject/iternal/service"
	"awesomeProject/iternal/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

// StudentAPI писать примерно так всегда (пока не пойму что так делать нельзя)
type StudentAPI struct {
}

func NewStudentAPI() *StudentAPI {
	return &StudentAPI{}
}

// получаю переменную сервиса (структуры service.StudentService) обЪект
var studentService = service.NewStudentService()

// Save метод обработчика для создания объекта domain.Student
func (sa *StudentAPI) Save(c *gin.Context) {

	// это пустая переменная которая копирует поля jsona из пришедшего запроса
	var body models.SaveStudent

	// тут происходит перенос данных из джейсона в переменную body
	if err := tools.RequestBinderBody(&body, c); err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	// вызывается метод сервиса для сохранения
	// и записывает результат его выполнения в переменную result
	result, err := studentService.Save(body)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	// тут передается ответ на запрос
	c.JSON(http.StatusCreated, result)
}

func (sa *StudentAPI) Get(c *gin.Context) {

	result, err := studentService.Get()

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (sa *StudentAPI) GetByID(c *gin.Context) {
	// id := c.Param("id") это является ключом id запроса
	id := c.Param("id")

	result, err := studentService.GetByID(id)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, result)
}
