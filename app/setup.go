package app

import (
	"github.com/aidenfine/go_carmeet_backend/config"
	"github.com/aidenfine/go_carmeet_backend/database"
	"github.com/aidenfine/go_carmeet_backend/router"
)

func SetupAndRunApp() error {
	err := config.LoadEnv()
	if err != nil {
		return err
	}

	// start the database
	db, err := database.ConnectToDatabase()
	if err != nil {
		return err
	}
	return router.StartServer(db)

}
