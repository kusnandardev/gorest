package main

import (
	"RestGo/pkg/infrastucture/db/inmemory"
	"RestGo/pkg/infrastucture/rest/router"
	"RestGo/pkg/infrastucture/simulate_log"
)

func main() {

	err := simulate_log.InitLogFile()
	if err != nil {
		panic(err)
	}

	inmemory.InitCache()
	routersInit := router.InitRouter()

	routersInit.Run(":8000")
}
