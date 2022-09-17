package main

import (
	"orba-back-end/Config"
	"orba-back-end/routes"
	"github.com/gin-gonic/gin"

)


func main() {
	r := gin.Default()

	Config.Connection()
	// routes.UserRoute(r)
	// routes.ProviderRoute(r)
	//  routes.ProuductRoute(r)
	//  routes.AdditionalRoute(r)
	 routes.CategoryRoute(r)
	//  routes.ReclamationRoute(r)
	//  routes.CategoryRoute(r)
	//  routes.CommandRoute(r)
	//  routes.Command_lineRoute(r)
	//  routes.LocationRoute(r)
	//  routes.FavoriteRoute(r)
	//  routes.TypeRoute(r)
    //  routes.RoteRoute(r)
	routes.UserRoute(r)


	r.Run(":8080")
	
}
