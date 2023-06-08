package actions

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uees/hidedomain/models"
	"github.com/uees/hidedomain/services"
)

func Login(c *gin.Context) {
	var form models.LoginForm
	if err := c.BindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"status":  "failed",
			"message": err.Error(),
			"data":    map[string]any{},
		})
		return
	}

	loginInfo := &models.LoginInfo{
		Username:  form.Username,
		Password:  form.Password,
		IP:        c.ClientIP(),
		UserAgent: c.GetHeader("User-Agent"),
	}

	if _, err := services.Login(loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"status":  "failed",
			"message": err.Error(),
			"data":    map[string]any{},
		})
		return
	}

	hours := 24 * time.Hour * time.Duration(1) // 24 h
	token, _ := services.GenerateJWT(form.Username, hours)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"message": "Get API token success.",
		"data": map[string]any{
			"token":              token,
			"refresh_expires_in": int64(hours / time.Second),
		},
	})
}

func Profile(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	body := gin.H{
		"success": true,
		"status":  "success",
		"message": "Get user profile success.",
		"data":    user,
	}

	c.JSON(http.StatusOK, body)
}
