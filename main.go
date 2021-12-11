package main

import (
	"keyValueProject/core/env"
	"keyValueProject/routers"
	"keyValueProject/service"
)

func main() {
	env.Init()
	service.Init()
	service.FetchValuesFromFile()
	service.StartInterval()

	router := routers.SetupRouter()

	router.Run()
}
