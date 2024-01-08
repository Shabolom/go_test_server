package main

import (
	"awesomeProject/config"
	"awesomeProject/iternal/routes"
)

func main() {

	// config.InitPgSQL инициализируем подключение к базе данных
	err := config.InitPgSQL()
	if err != nil {
		panic(err)
	}

	// конфигурация (инициализация) end point или ручка (можно назвать имя запроса)
	// (как api student) URLов пример (localhost, 8080) конфигурация всех URLов которые будет
	// обрабатывать сервер

	r := routes.SetupRouter()

	// запуск сервера
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
