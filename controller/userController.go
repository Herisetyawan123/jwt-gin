package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heirsetyawan233/jwt-gin/initializers"
	"github.com/heirsetyawan233/jwt-gin/models"
	"github.com/heirsetyawan233/jwt-gin/request"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// ambil email dan pass dari req body
	if c.Bind(&request.SignUpBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.SignUpBody.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
	}

	user := models.User{Email: request.SignUpBody.Email, Password: string(hash), Name: request.SignUpBody.Name}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your success to signup",
	})
}
