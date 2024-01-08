package routes

import (
	"awesomeProject/iternal/api"
	"awesomeProject/iternal/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRouter производим подключение к библиотеке gin (фреймворку)
func SetupRouter() *gin.Engine {
	// переменная r предоставляет функционал работы с сетью
	r := gin.Default()

	// получение переменной с функциями-обработчиками
	student := api.NewStudentAPI()
	teacher := api.NewTeacherAPI()
	user := api.NewUserAPI()

	r.POST("/api/user/register", user.Save)
	r.POST("/api/user/login", middleware.Passport().LoginHandler)

	authRequired := r.Group("/")
	authRequired.Use(middleware.JwtAuthMiddleware())

	{
		authRequired.POST("/api/student", student.Save)
		authRequired.GET("/api/student", student.Get)
		authRequired.GET("/api/student/:id", student.GetByID)

		authRequired.POST("/api/teacher", teacher.Save)
		authRequired.GET("/api/teacher", teacher.Get)
		authRequired.GET("/api/teacher/:id", teacher.GetByID)

		authRequired.POST("/api/user", user.Save)
		authRequired.GET("/api/user", user.Get)
		authRequired.GET("/api/user/:id", user.GetByKey)
		authRequired.PUT("/api/user/:id", user.Update)
	}
	// описание URLа где POST это метод запроса, /api/student имя запроса,
	//student.Save это функция-обработчик (это функция при обращении на данный запрос)

	return r
}
