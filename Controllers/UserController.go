package controllers

import (
	 "orba-back-end/entities"
	 "orba-back-end/Config"


	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	 users:= []entities.User{}
	Config.DB.Find(&users)
	 c.JSON(200,&users)
    c.String(200,"hello")
}
 func CreateUser(c *gin.Context){
 	var user entities.User
 	c.BindJSON(&user)
	Config.DB.Create(&user)
 	c.JSON(200,&user)
 }