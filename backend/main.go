package main

import (
	"github.com/kimnopal/maya/config"
)

func main() {
	app := config.NewFiber()
	db := config.NewDatabase()
	validate := config.NewValidator()

	config.Bootstrap(&config.BootstrapConfig{
		App:      app,
		DB:       db,
		Validate: validate,
	})

	err := app.Listen("localhost:8000")
	if err != nil {
		panic(err)
	}
}

// migrate -database "mysql://root@tcp(localhost:3306)/maya" -path db/migrations up
