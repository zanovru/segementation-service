package main

import (
	"segmenatationService/internal/app"
)

const configPath = "config/config.yaml"

func main() {
	app.Run(configPath)

}
