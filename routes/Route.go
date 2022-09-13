package routes

import (
	controllers "orba-back-end/Controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine){ 
	router.GET("/", controllers.GetUsers)
	 router.POST("/Post",controllers.CreateUser)
}