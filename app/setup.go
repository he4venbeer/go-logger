package app

import (
	"go-logger/config"
	"go-logger/database"
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

	return nil
}
