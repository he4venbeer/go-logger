package app

import (
	"go-logger/config"
	"go-logger/database"
	"go-logger/routes"
	"os"

	"github.com/gofiber/fiber/v2"
)

func SetupAndRunApp() error {
	// load ENV
	err := config.LoadENV()
	if err != nil {
		return err
	}

	// start database
	err = database.StartMongoDB()
	if err != nil {
		return err
	}

	// defer closing database
	defer database.CloseMongoDB()

	// create app
	app := fiber.New()

	// set up routes
	routes.SetupRoutes(app)

	port := os.Getenv("APP_PORT")
	app.Listen(":" + port)

	return nil
}
