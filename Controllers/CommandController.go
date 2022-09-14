package controllers

import (
	 "orba-back-end/entities"
	 "orba-back-end/Config"
	 "net/http"

	"github.com/gin-gonic/gin"
)

func GetCommand_(c *gin.Context) {
	var command []entities.Command
	Config.DB.Find(&command)
	c.JSON(http.StatusOK, gin.H{"data": command})
}
 func CreateCommand(c *gin.Context){
 	var command entities.Command
 	c.BindJSON(&command)
	Config.DB.Create(&command)
 	c.JSON(200,&command)
 }
 func DeleteCommand(c *gin.Context){
	var command entities.Command 
	Config.DB.Where("ID=?",c.Param("ID")).Delete(&command)
	c.JSON(200,&command)
 }
 func UpdateCommand(c *gin.Context){
	var command entities.Command 
	Config.DB.Where("ID=?",c.Param("ID")).First(&command)
	c.BindJSON(&command)
	Config.DB.Save(&command)
	c.JSON(200,&command)

 }
