package main

import (
	// "fmt"
	"rooster-blog/router"
)

func main() {
	router := router.InitRouter()
	router.Run(":9100")
}