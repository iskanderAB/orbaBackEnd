package routes

import (
	controllers "orba-back-end/Controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine){ 
	 router.GET("/users", controllers.GetUsers)
	 router.POST("/users",controllers.CreateUser)
	 router.DELETE("/users/:ID",controllers.DeleteUser)
	 router.PUT("/users/:ID",controllers.UpdateUser)

}
func ProviderRoute(router *gin.Engine){ 
	router.GET("/Providers", controllers.GetProviders)
	router.POST("/Providers",controllers.CreateProvider)
	router.DELETE("/Providers/:ID",controllers.DeleteProvider)
	router.PUT("/PRoviders/:ID",controllers.UpdateProvider)

}