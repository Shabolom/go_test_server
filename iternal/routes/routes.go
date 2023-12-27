package routes

import (
	"awesomeProject/iternal/api"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	student := api.NewStudentAPI()

	r.POST("/api/student", student.Save)

	return r
}
