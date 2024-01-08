package api

import (
	"awesomeProject/iternal/models"
	"awesomeProject/iternal/service"
	"awesomeProject/iternal/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserAPI struct {
}

func NewUserAPI() *UserAPI {
	return &UserAPI{}
}

var userService = service.NewUserService()

func (ua *UserAPI) Save(c *gin.Context) {

	var body models.SaveUser

	if err := tools.RequestBinderBody(&body, c); err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	result, err := userService.Save(body)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (ua *UserAPI) Get(c *gin.Context) {

	result, err := userService.Get()

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (ua *UserAPI) GetByKey(c *gin.Context) {

	id := c.Param("id")

	result, err := userService.GetByKey(id)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (ua *UserAPI) Update(c *gin.Context) {

	id := c.Param("id")

	var body models.SaveUser

	if err := tools.RequestBinderBody(&body, c); err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	result, err := userService.Update(id, body)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, result)
}
