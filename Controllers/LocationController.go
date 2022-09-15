package controllers

import (
	 "orba-back-end/entities"
	 "orba-back-end/Config"
	 "net/http"

	"github.com/gin-gonic/gin"
)

func GetLocation(c *gin.Context) {
	var locs []entities.Location
	Config.DB.Find(&locs)
	c.JSON(http.StatusOK, gin.H{"LOCATIONS": locs})
}
 func CreateLocation(c *gin.Context){
 	var loc entities.Location
 	c.BindJSON(&loc)
	Config.DB.Create(&loc)
 	c.JSON(200,&loc)
 }
 func DeleteLocation(c *gin.Context){
	var loc entities.Location 
	Config.DB.Where("ID=?",c.Param("ID")).Delete(&loc)
	c.JSON(200,&loc)
 }
 func UpdateLocation(c *gin.Context){
	var loc entities.Location
	Config.DB.Where("ID=?",c.Param("ID")).First(&loc)
	c.BindJSON(&loc)
	Config.DB.Save(&loc)
	c.JSON(200,&loc)

 }
