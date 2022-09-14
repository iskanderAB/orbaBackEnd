package main

import (
	"orba-back-end/Config"
	"orba-back-end/routes"
	"github.com/gin-gonic/gin"

)


func main() {
	r := gin.Default()

	Config.Connection()
	routes.UserRoute(r)

	r.Run(":8080")
	
}
