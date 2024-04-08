package routes

import (
	"awesomeProject/iternal/api"
	"awesomeProject/iternal/middleware"
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// SetupRouter производим подключение к библиотеке gin (фреймворку)
func SetupRouter() *gin.Engine {
	// переменная r предоставляет функционал работы с сетью
	r := gin.New()

	// получение переменной с функциями-обработчиками
	student := api.NewStudentAPI()
	teacher := api.NewTeacherAPI()
	user := api.NewUserAPI()

	r.POST("/api/user/register", user.Save)
	r.POST("/api/user/login", middleware.Passport().LoginHandler)
	r.DELETE("/api/user/:key/:value", user.Delet)
	//r.DELETE("/api/student", student.Delet)
	//r.DELETE("/api/teacher", teacher.Delet)

	authRequired := r.Group("/")
	authRequired.Use(middleware.Passport().MiddlewareFunc())

	{
		authRequired.POST("/api/student", student.Save)
		authRequired.GET("/api/student", student.Get)
		authRequired.GET("/api/student/:id", student.GetByID)

		authRequired.POST("/api/teacher", teacher.Save)
		authRequired.GET("/api/teacher", teacher.Get)
		authRequired.GET("/api/teacher/:id", teacher.GetByID)

		authRequired.GET("/api/user", user.Get)
		authRequired.GET("/api/user/:key", user.GetByKey)
		authRequired.PUT("/api/user/:id", user.Update)
		r.POST("/", func(c *gin.Context) {
			values := "https://zetcode.com/golang/getpostrequest123"
			jsonData := []byte(values)

			var buf bytes.Buffer

			g := gzip.NewWriter(&buf)

			if _, err := g.Write(jsonData); err != nil {
				println(err)
				return
			}
			if err := g.Close(); err != nil {
				println(err)
				return
			}

			client := &http.Client{}

			req, err := http.NewRequest("POST", "http://localhost:8000/", &buf)

			if err != nil {
				fmt.Println(err)
				return
			}

			// хэдер обозначающий что данные будут сжаты
			req.Header.Add("Content-Encoding", `gzip`)
			req.Header.Add("Accept-Encoding", `gzip`)

			resp, err := client.Do(req)

			defer resp.Body.Close()

			respBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			respString := string(respBytes)

			fmt.Printf(respString)
		})
	}
	// описание URLа где POST это метод запроса, /api/student имя запроса,
	//student.Save это функция-обработчик (это функция при обращении на данный запрос)

	return r
}
