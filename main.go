package main

import (
	"rooster-blog/router"

)

func main() {
	router := router.InitRouter()
	router.Run(":9100")
}