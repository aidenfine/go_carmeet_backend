package main

import (
	"github.com/aidenfine/go_carmeet_backend/app"
)

func main() {
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
