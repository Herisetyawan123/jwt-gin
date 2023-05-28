package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

func SignIn(c *gin.Context) {
	var user models.User

	// ambil email dan pass dari req body
	if c.Bind(&request.SignInBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	initializers.DB.First(&user, "email = ?", request.SignInBody.Email)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"errors": "Email Belum Terdaftar",
		})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.SignInBody.Password)) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password yang anda masukan salah",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Terjadi kesalahan ketika menerjemahkan token",
		})
	}

	// cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Anda berhasil login",
		"token":   tokenString,
	})

}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "I'am login",
	})
}
