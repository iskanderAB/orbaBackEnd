package controllers

import (
	"orba-back-end/entities"

	"github.com/gin-gonic/gin"
	"orba-back-end/Config"
)

func GetAdditional(c *gin.Context){
	var additional []entities.Additional
	Config.DB.Find(&additional)
	c.JSON(200,additional)


}
func UpdateAdditonal(c *gin.Context){
	var additional entities.Additional
	Config.DB.Where("ID",c.Param("ID")).First(&additional)
	c.BindJSON(&additional)
	Config.DB.Save(&additional)
	c.JSON(200,&additional)
}
func CreateAdditional(c *gin.Context){
	var additional entities.Additional
	c.BindJSON(&additional)
	Config.DB.Create(&additional)
	c.JSON(200,&additional)
}
func DeleteAdditonal (c *gin.Context){
	var additional entities.Additional
	Config.DB.Where("ID",c.Param("ID")).Delete(&additional)
	c.JSON(200,&additional)

}