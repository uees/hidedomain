package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uees/hidedomain/models"
	"github.com/uees/hidedomain/services"
)

func ShowList(c *gin.Context) {
	domain := c.Param("domain")
	var wlists []models.Whitelist
	if err := services.GetWhiteListByDomain(domain, &wlists); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"data":    wlists,
	})
}

func ClearList(c *gin.Context) {
	domain := c.Param("domain")
	if err := services.ClearWhiteListByDomain(domain); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"message": "clear all rules",
	})
}

func AddIPRule(c *gin.Context) {
	domain := c.Param("domain")
	var rule models.RuleForm
	if err := c.BindJSON(&rule); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := services.AddIPRule(domain, &rule); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"message": "add rule success",
	})
}

func UpdateIPRule(c *gin.Context) {
	ruleid := c.Param("ruleid")
	var rule models.RuleForm
	if err := c.BindJSON(&rule); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := services.UpdateIPRule(ruleid, &rule); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"message": "update rule success",
	})
}

func DeleteIPRule(c *gin.Context) {
	ruleid := c.Param("ruleid")

	if err := services.DeleteIPRule(ruleid); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"message": "delete rule success",
	})
}
