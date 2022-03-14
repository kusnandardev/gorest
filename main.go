package main

import (
	"RestGo/pkg/infrastucture/db/cache"
	"RestGo/pkg/infrastucture/rest/router"
	"RestGo/pkg/infrastucture/simulate_log"
)

func main() {

	err := simulate_log.InitLogFile()
	if err != nil {
		panic(err)
	}
	c := cache.NewCache()
	c.SetCache()
	routersInit := router.InitRouter()

	routersInit.Run(":8000")
}
