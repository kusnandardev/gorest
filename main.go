package main

import (
	"RestGo/pkg/infrastucture/rest/router"
)

func main() {
	routersInit := router.InitRouter()

	routersInit.Run(":8000")
}
