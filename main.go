package main

import (
	"github.com/ValentinAltamirano1/WashUp-Api/database"
	"github.com/ValentinAltamirano1/WashUp-Api/serve"
)

func main() {
	err := Serve()

	if err != nil {
		panic(err)
	}
}

func Serve() error {
	app := serve.SetRouters()

	database.Connect()

	return app.Listen(":3000")	
}

