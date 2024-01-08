package api

import (
	"awesomeProject/iternal/models"
	"awesomeProject/iternal/service"
	"awesomeProject/iternal/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TeacherAPI struct {
}

func NewTeacherAPI() *TeacherAPI {
	return &TeacherAPI{}
}

var teacherService = service.NewTeacherService()

func (ta TeacherAPI) Save(c *gin.Context) {
	var body models.SaveTeacher

	if err := tools.RequestBinderBody(&body, c); err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	result, err := teacherService.Save(body)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (ta TeacherAPI) Get(c *gin.Context) {

	result, err := teacherService.Get()

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (ta TeacherAPI) GetByID(c *gin.Context) {

	id := c.Param("id")
	result, err := teacherService.GetByID(id)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, result)
}
