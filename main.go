package main

import (
	"keyValueProject/core/env"
	"keyValueProject/core/logging"
	"keyValueProject/routers"
	"keyValueProject/service"
)

func main() {
	env.Init()
	service.Init()
	service.FetchValuesFromFile()
	service.StartInterval()
	logging.Init()

	router := routers.SetupRouter()

	router.Run()
}
