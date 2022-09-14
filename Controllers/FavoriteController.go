package controllers

import (
	 "orba-back-end/entities"
	 "orba-back-end/Config"
	 "net/http"

	"github.com/gin-gonic/gin"
)

func GetFaviortes(c *gin.Context) {
	var favorites []entities.Favorite
	Config.DB.Find(&favorites)
	c.JSON(http.StatusOK, gin.H{"data": favorites})
}
 func CreateFavorite(c *gin.Context){
 	var favorite entities.Favorite
 	c.BindJSON(&favorite)
	Config.DB.Create(&favorite)
 	c.JSON(200,&favorite)
 }
 func DeleteFavorite(c *gin.Context){
	var favorite entities.Favorite 
	Config.DB.Where("ID=?",c.Param("ID")).Delete(&favorite)
	c.JSON(200,&favorite)
 }
 func UpdateFavorite(c *gin.Context){
	var favorite entities.Favorite 
	Config.DB.Where("ID=?",c.Param("ID")).First(&favorite)
	c.BindJSON(&favorite)
	Config.DB.Save(&favorite)
	c.JSON(200,&favorite)

 }
