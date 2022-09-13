package main

import (
	"orba-back-end/Config"
	"orba-back-end/routes"
	"github.com/gin-gonic/gin"

)


func main() {

	router:=gin.New()
	Config.Connection()
	routes.UserRoute(router)

	router.Run(":8080")
	
}
