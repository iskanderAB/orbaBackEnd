package controllers

import (
	 "orba-back-end/entities"
	 "orba-back-end/Config"
	 "net/http"

	"github.com/gin-gonic/gin"
)

func GetProviders(c *gin.Context) {
	var provs []entities.Provider
	Config.DB.Find(&provs)
	c.JSON(http.StatusOK, gin.H{"data": provs})
}
 func CreateProvider(c *gin.Context){
 	var prov entities.Provider
 	c.BindJSON(&prov)
	Config.DB.Create(&prov)
 	c.JSON(200,&prov)
 }
 func DeleteProvider(c *gin.Context){
	var prov entities.Provider 
	Config.DB.Where("ID=?",c.Param("ID")).Delete(&prov)
	c.JSON(200,&prov)
 }
 func UpdateProvider(c *gin.Context){
	var prov entities.Provider 
	Config.DB.Where("ID=?",c.Param("ID")).First(&prov)
	c.BindJSON(&prov)
	Config.DB.Save(&prov)
	c.JSON(200,&prov)

 }
