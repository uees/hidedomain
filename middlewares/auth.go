package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4/request"
	"github.com/uees/hidedomain/models"
	"github.com/uees/hidedomain/services"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := request.OAuth2Extractor.ExtractToken(c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "unauthorized",
				"message": err.Error(),
			})
			return
		}

		claims, err := services.ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "unauthorized",
				"message": err.Error(),
			})
			return
		}

		var user models.User
		result := services.DB.Where("username = ?", claims.User).First(&user)
		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "unauthorized",
				"message": err.Error(),
			})
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
