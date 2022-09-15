package controllers

import (
	 "orba-back-end/entities"
	 "orba-back-end/Config"
	 "net/http"

	"github.com/gin-gonic/gin"
)

func GetCommand_lines(c *gin.Context) {
	var cl []entities.User
	Config.DB.Find(&cl)
	c.JSON(http.StatusOK, gin.H{"COMMAND_LINES": cl})
}
 func CreateCommand_line(c *gin.Context){
 	var cl entities.Command_line
 	c.BindJSON(&cl)
	Config.DB.Create(&cl)
 	c.JSON(200,&cl)
 }
 func DeleteCommand_line(c *gin.Context){
	var cl entities.Command_line 
	Config.DB.Where("ID=?",c.Param("ID")).Delete(&cl)
	c.JSON(200,&cl)
 }
 func UpdateCommand_line(c *gin.Context){
	var cl entities.Command_line 
	Config.DB.Where("ID=?",c.Param("ID")).First(&cl)
	c.BindJSON(&cl)
	Config.DB.Save(&cl)
	c.JSON(200,&cl)

 }