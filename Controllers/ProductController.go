package controllers

import (
	 "orba-back-end/entities"
	 "orba-back-end/Config"
	 "net/http"

	"github.com/gin-gonic/gin"
)

func GetProuduct(c *gin.Context) {
	var products []entities.Product
	Config.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"PRODUCTS": products})
}
 func CreateProduct(c *gin.Context){
 	var prod entities.Product
 	c.BindJSON(&prod)
	Config.DB.Create(&prod)
 	c.JSON(200,&prod)
 }
 func DeleteProduct(c *gin.Context){
	var Prod entities.Product 
	Config.DB.Where("ID=?",c.Param("ID")).Delete(&Prod)
	c.JSON(200,&Prod)
 }
 func UpdateProduct(c *gin.Context){
	var prod entities.Product
	Config.DB.Where("ID=?",c.Param("ID")).First(&prod)
	c.BindJSON(&prod)
	Config.DB.Save(&prod)
	c.JSON(200,&prod)

 }
