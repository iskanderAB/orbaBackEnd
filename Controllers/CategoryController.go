package controllers

import (
	"orba-back-end/entities"

	"github.com/gin-gonic/gin"
	"orba-back-end/Config"
)

func GetCategorys(c *gin.Context){
	var category []entities.Category
	Config.DB.Find(&category)
	c.JSON(200,&category)


}
func UpdateCategory(c *gin.Context){
	var category entities.Additional
	Config.DB.Where("ID",c.Param("ID")).First(&category)
	c.BindJSON(&category)
	Config.DB.Save(&category)
	c.JSON(200,&category)
}
func CreateCategory(c *gin.Context){
	var category entities.Category
	c.BindJSON(&category)
	Config.DB.Create(&category)
	c.JSON(200,&category)
}
func DeleteCategory (c *gin.Context){
	var category entities.Category
	Config.DB.Where("ID",c.Param("ID")).Delete(&category)
	c.JSON(200,&category)

}