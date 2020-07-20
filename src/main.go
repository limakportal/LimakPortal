package main

import (
	"limakcv/src/app"
	"limakcv/src/config"
	"os"
)

func main() {
	config := config.GetConfig()
	port := os.Getenv("PORT")
	app := &app.App{}
	app.Initialize(config)
	app.Run(":" + port)

}
