package main

import (
	"go_simple_mvc/routers"
)

func main() {
	router := routers.SetupRouter()
	router.Run()
}
