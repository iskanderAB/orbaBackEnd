package controllers

import (
	 "orba-back-end/entities"
	 "orba-back-end/Config"
	 "net/http"

	"github.com/gin-gonic/gin"
)

func GetNotification(c *gin.Context) {
	var users []entities.User
	Config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"NOTIFICATIONS": users})
}
 func CreateNotification(c *gin.Context){
 	var not entities.Notification
 	c.BindJSON(&not)
	Config.DB.Create(&not)
 	c.JSON(200,&not)
 }
 func DeleteNotification(c *gin.Context){
	var not entities.Notification 
	Config.DB.Where("ID=?",c.Param("ID")).Delete(&not)
	c.JSON(200,&not)
 }
 func UpdateNotification(c *gin.Context){
	var not entities.Notification 
	Config.DB.Where("ID=?",c.Param("ID")).First(&not)
	c.BindJSON(&not)
	Config.DB.Save(&not)
	c.JSON(200,&not)

 }
