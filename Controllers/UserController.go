package controllers

import (
	"net/http"
	"orba-back-end/Config"
	"orba-back-end/entities"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUsers(c *gin.Context) {
	var users []entities.User
	Config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"USERS": users})
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
 func SingUp(c*gin.Context){
	//Get the email and the password
	var body struct{
		Fullname string
		Role string
		Gendre string
		Email string
		Password string
		Telephone int32
		
	}
	if(c.BindJSON(&body) != nil){
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"failed to read body"})
		return
		}
		//Hash the password
		hash,err:=bcrypt.GenerateFromPassword([]byte(body.Password),10)
		if err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":"failed to hash password"})
		return
		}
		//create user
		user:=entities.User{
			Model:         gorm.Model{},
			Fullname:      body.Fullname,
			Role:          body.Role,
			Email:         body.Email,
			Gender:        body.Gendre,
			Telephone:     body.Telephone,
			Password:      string(hash),
			Providers:     []entities.Provider{},
			Favorites:     []entities.Favorite{},
			Reclamations:  []entities.Reclamation{},
			Commands:      []entities.Command{},
			Rotes:         []entities.Rote{},
			Notifications: []entities.Notification{},
		} 
		result:=Config.DB.Create(&user)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":"failed to create the user"})
		return 
	    }
	//respond
	c.JSON(http.StatusAccepted,gin.H{})
}
func Login(c*gin.Context){
	//Get the email and the password
	var body struct{
	
		Email string
		Password string
		
		
	}
	if(c.BindJSON(&body) != nil){
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"failed to read body"})
		return
		}
		// look up for informations
		var user entities.User
		Config.DB.First(&user,"email=?",body.Email)
		if user.ID==0 {c.JSON(http.StatusBadRequest,gin.H{
			"err":"Invalid email or Password"})
			return
		}
		//Compare passwords 
		err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(body.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":"invalid password"})
			return

		}
		//generate a jwt token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub":user.ID,
			"exp": time.Now().Add(time.Hour * 24 * 30 ).Unix(),
		})
		
		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
		if err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":"failed to get token",
		})
		return	
		}
		//send it back 
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorisation",tokenString,3600 *24 *30,"","" ,false,true)
		c.JSON(http.StatusAccepted,gin.H{})
	
		

}


