package api

import (
	"awesomeProject/iternal/models"
	"awesomeProject/iternal/service"
	"awesomeProject/iternal/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StudentAPI struct {
}

func NewStudentAPI() *StudentAPI { return &StudentAPI{} }

var studentService = service.NewStudentService()

func (sa StudentAPI) Save(c *gin.Context) {

	var body models.SaveStudent

	if err := tools.RequestBinderBody(&body, c); err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	result, err := studentService.Save(body)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}
	c.JSON(http.StatusCreated, result)
}
