package authcontroller

import (
	// "encoding/json"
	"net/http"
	"time"

	"github.com/balada-raja/GET/initializers"
	"github.com/balada-raja/GET/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	// input email and password
	var input struct {
		Email    string
		Password string
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//data user berdasarkan email
	var user models.Users
	if err := initializers.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "email atau password salah"}
			c.JSON(http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	}

	// cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		response := map[string]string{"message": "email atau password salah"}
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	//creat jwt
	expTime := time.Now().Add(time.Minute * 1)
	claims := &initializers.JWTClaim{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	//
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//signed token
	token, err := tokenAlgo.SignedString(initializers.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": "email atau password salah"}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//set token ke cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})


	response := map[string]string{"message": "Login berhasil"}
	c.JSON(http.StatusOK, response)

}

func Register(c *gin.Context) {
	// input json
	var userInput models.Users

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// hash pass bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	//insert
	initializers.DB.Create(&userInput)
	c.JSON(http.StatusOK, gin.H{"users": "success"})

}

func Logout(c *gin.Context) {
	//hapus token
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "Logout berhasil"}
	c.JSON(http.StatusOK, response)
}
