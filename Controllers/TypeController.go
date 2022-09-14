package controllers

import (
	 "orba-back-end/entities"
	 "orba-back-end/Config"
	 "net/http"

	"github.com/gin-gonic/gin"
)

func GetType(c *gin.Context) {
	var tps []entities.Type
	Config.DB.Find(&tps)
	c.JSON(http.StatusOK, gin.H{"data": tps})
}
 func CreateType(c *gin.Context){
 	var tp entities.Type
 	c.BindJSON(&tp)
	Config.DB.Create(&tp)
 	c.JSON(200,&tp)
 }
 func Delete(c *gin.Context){
	var tp entities.Type 
	Config.DB.Where("ID=?",c.Param("ID")).Delete(&tp)
	c.JSON(200,&tp)
 }
 func UpdateType(c *gin.Context){
	var tp entities.Type 
	Config.DB.Where("ID=?",c.Param("ID")).First(&tp)
	c.BindJSON(&tp)
	Config.DB.Save(&tp)
	c.JSON(200,&tp)

 }
