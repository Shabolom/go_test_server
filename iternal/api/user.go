package api

// слой работы с запросом
import (
	"awesomeProject/iternal/models"
	"awesomeProject/iternal/service"
	"awesomeProject/iternal/tools"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
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

	// c.Request.Body происходит получение тела от клиента поля Body который на даееый момент
	// является reader ом (это оболчка в которую оборачивается байты) и чтобы прочитать
	// контент который хранится в body нужно ссначало распакавать reader и для этого мы
	// используем функцию io.ReadAll которая возвращает нам ошибку и байты которые мы потом
	// де сериализируем (преобразуем из байтового формата в объект)
	gog, err := io.ReadAll(c.Request.Body)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	// тут происходит процесс де сериализации (преобразуем из байтового формата в объект)
	err = json.Unmarshal(gog, &body)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	defer c.Request.Body.Close()

	//if err := tools.RequestBinderBody(&body, c); err != nil {
	//	tools.CreateError(http.StatusBadRequest, err, c)
	//	return
	//}

	result, err := userService.Save(body)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}
	fmt.Println(c.Request.Header)

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

	id := c.Param("key")

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

func (ua *UserAPI) Delet(c *gin.Context) {

	id := c.Param("key")
	serch := c.Param("value")
	result, err := userService.Delet(id, serch)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}
	fmt.Println(c.Request.Header)
	c.JSON(http.StatusCreated, result)
}
