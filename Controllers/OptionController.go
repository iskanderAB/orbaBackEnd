package controllers

import (
	 "orba-back-end/entities"
	 "orba-back-end/Config"
	 "net/http"

	"github.com/gin-gonic/gin"
)

func GetOptions(c *gin.Context) {
	var opts []entities.Option
	Config.DB.Find(&opts)
	c.JSON(http.StatusOK, gin.H{"OPTIONS": opts})
}
 func CreateOption(c *gin.Context){
 	var opt entities.Option
 	c.BindJSON(&opt)
	Config.DB.Create(&opt)
 	c.JSON(200,&opt)
 }
 func DeleteOption(c *gin.Context){
	var opt entities.Option 
	Config.DB.Where("ID=?",c.Param("ID")).Delete(&opt)
	c.JSON(200,&opt)
 }
 func UpdateOption(c *gin.Context){
	var opt entities.Option
	Config.DB.Where("ID=?",c.Param("ID")).First(&opt)
	c.BindJSON(&opt)
	Config.DB.Save(&opt)
	c.JSON(200,&opt)

 }
