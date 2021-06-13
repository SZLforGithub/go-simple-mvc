package main

import (
	"github.com/SZLforGithub/go-simple-mvc/routers"
)

func main() {
	router := routers.SetupRouter()
	router.Run()
}
