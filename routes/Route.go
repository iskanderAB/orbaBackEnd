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
func AdditionalRoute(router *gin.Engine){
	router.GET("/Additionals",controllers.GetAdditional)
	router.POST("/Additionals",controllers.CreateAdditional)
	router.DELETE("/Additionals/:ID",controllers.DeleteAdditonal)
	router.PUT("/Additionals/:ID",controllers.UpdateAdditonal)

}
func CategoryRoute(router *gin.Engine){
	router.GET("/Categorys",controllers.GetCategorys)
	router.POST("/Categorys",controllers.CreateCategory)
	router.DELETE("/Categorys/:ID",controllers.DeleteCategory)
	router.PUT("/Categorys/:ID",controllers.UpdateCategory)

}
func Command_lineRoute(router *gin.Engine){
	router.GET("/Command_lines",controllers.GetCommand_lines)
	router.POST("/Command_lines",controllers.CreateCommand_line)
	router.DELETE("/Command_lines/:ID",controllers.DeleteCommand_line)
	router.PUT("/Command_lines/:ID",controllers.UpdateCommand_line)

}
func CommandRoute(router *gin.Engine){
	router.GET("/Commands",controllers.GetCommands)
	router.POST("/Commands",controllers.CreateCommand)
	router.DELETE("/Commands/:ID",controllers.DeleteCommand)
	router.PUT("/Commands/:ID",controllers.UpdateCommand)

}
func FavoriteRoute(router *gin.Engine){
	router.GET("/Favorites",controllers.GetFaviortes)
	router.POST("/Favorites",controllers.CreateFavorite)
	router.DELETE("/Favorites/:ID",controllers.DeleteFavorite)
	router.PUT("/Favorites/:ID",controllers.UpdateFavorite)

}
func LocationRoute(router *gin.Engine){
	router.GET("/Locations",controllers.GetLocation)
	router.POST("/Locations",controllers.CreateLocation)
	router.DELETE("/Locations/:ID",controllers.DeleteLocation)
	router.PUT("/Locations/:ID",controllers.UpdateLocation)

}
func NotificationRoute(router *gin.Engine){
	router.GET("/Notifications",controllers.GetNotification)
	router.POST("/Notifications",controllers.CreateNotification)
	router.DELETE("/Notifications/:ID",controllers.DeleteNotification)
	router.PUT("/Notifications/:ID",controllers.UpdateNotification)

}
func OptionRoute(router *gin.Engine){
	router.GET("/Options",controllers.GetOptions)
	router.POST("/Options",controllers.CreateOption)
	router.DELETE("/Options/:ID",controllers.CreateOption)
	router.PUT("/Options/:ID",controllers.UpdateOption)

}
func ProuductRoute(router *gin.Engine){
	router.GET("/Products",controllers.GetProuduct)
	router.POST("/Products",controllers.CreateProduct)
	router.DELETE("/Products/:ID",controllers.DeleteProduct)
	router.PUT("/Products/:ID",controllers.UpdateProduct)

}
func ReclamationRoute(router *gin.Engine){
	router.GET("/Reclamations",controllers.GetReclamations)
	router.POST("/Reclamations",controllers.CreateReclamation)
	router.DELETE("/Reclamations/:ID",controllers.DeleteReclamtion)
	router.PUT("/Reclamations/:ID",controllers.UpdateReclamation)

}
func RoteRoute(router *gin.Engine){
	router.GET("/Rotes",controllers.GetRotes)
	router.POST("/Rotes",controllers.CreateRote)
	router.DELETE("/Rotes/:ID",controllers.DeleteRote)
	router.PUT("/Rotes/:ID",controllers.UpdateRote)

}
func TypeRoute(router *gin.Engine){
	router.GET("/Types",controllers.GetTypes)
	router.POST("/Types",controllers.CreateType)
	router.DELETE("/Types/:ID",controllers.DeleteType)
	router.PUT("/Types/:ID",controllers.UpdateType)

}