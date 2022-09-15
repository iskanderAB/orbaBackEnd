package controllers

import (
	 "orba-back-end/entities"
	 "orba-back-end/Config"
	 "net/http"

	"github.com/gin-gonic/gin"
)

func GetReclamations(c *gin.Context) {
	var recs []entities.Reclamation
	Config.DB.Find(&recs)
	c.JSON(http.StatusOK, gin.H{"RECLAMATIONS": recs})
}
 func CreateReclamation(c *gin.Context){
 	var rec entities.Reclamation
 	c.BindJSON(&rec)
	Config.DB.Create(&rec)
 	c.JSON(200,&rec)
 }
 func DeleteReclamtion(c *gin.Context){
	var rec entities.Reclamation 
	Config.DB.Where("ID=?",c.Param("ID")).Delete(&rec)
	c.JSON(200,&rec)
 }
 func UpdateReclamation(c *gin.Context){
	var rec entities.Reclamation 
	Config.DB.Where("ID=?",c.Param("ID")).First(&rec)
	c.BindJSON(&rec)
	Config.DB.Save(&rec)
	c.JSON(200,&rec)

 }
