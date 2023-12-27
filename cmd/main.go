package main

import (
	"awesomeProject/config"
	"awesomeProject/iternal/routes"
)

func main() {

	err := config.InitPgSQL()
	if err != nil {
		panic(err)
	}

	r := routes.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
