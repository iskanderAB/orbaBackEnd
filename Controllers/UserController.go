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
