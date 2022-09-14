package controllers

import (
	 "orba-back-end/entities"
	 "orba-back-end/Config"
	 "net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []entities.User
	Config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}
 func CreateUser(c *gin.Context){
 	var user entities.User
 	c.BindJSON(&user)
	Config.DB.Create(&user)
 	c.JSON(200,&user)
 }
 func DeleteUser(c *gin.Context){
	var user entities.User 
	Config.DB.Where("ID=?",c.Param("ID")).Delete(&user)
	c.JSON(200,&user)
 }
 func UpdateUser(c *gin.Context){
	var user entities.User 
	Config.DB.Where("ID=?",c.Param("ID")).First(&user)
	c.BindJSON(&user)
	Config.DB.Save(&user)
	c.JSON(200,&user)

 }
