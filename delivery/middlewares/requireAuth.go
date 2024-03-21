package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/balada-raja/GET/models"
	"github.com/balada-raja/GET/repository/initializers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	bearer := c.GetHeader("Authorization")
	if bearer == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "empty token"})
		return
	}

	tokenString := strings.Split(bearer, " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(initializers.JWT_KEY), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//find the user with token sub
		var user models.Users
		initializers.DB.First(&user, claims["sub"])

		if user.Id == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//go with request
		c.Set("user", user)
		c.Next()
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid token",
		})
	}
}