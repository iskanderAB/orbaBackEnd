package controllers

import (
	 "orba-back-end/entities"
	 "orba-back-end/Config"
	 "net/http"

	"github.com/gin-gonic/gin"
)

func GetRotes(c *gin.Context) {
	var rotes []entities.Rote
	Config.DB.Find(&rotes)
	c.JSON(http.StatusOK, gin.H{"ROTES": rotes})
}
 func CreateRote(c *gin.Context){
 	var rote entities.Rote
 	c.BindJSON(&rote)
	Config.DB.Create(&rote)
 	c.JSON(200,&rote)
 }
 func DeleteRote(c *gin.Context){
	var rote entities.Rote 
	Config.DB.Where("ID=?",c.Param("ID")).Delete(&rote)
	c.JSON(200,&rote)
 }
 func UpdateRote(c *gin.Context){
	var rote entities.Rote 
	Config.DB.Where("ID=?",c.Param("ID")).First(&rote)
	c.BindJSON(&rote)
	Config.DB.Save(&rote)
	c.JSON(200,&rote)

 }
