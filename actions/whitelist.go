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

// 清理ip列表，并清除iptables ip 列表
func ClearList(c *gin.Context) {
	domain := c.Param("domain")
	var wlists []models.Whitelist

	if err := services.GetWhiteListByDomain(domain, &wlists); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := services.ClearWhiteListByDomain(domain); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// clear all ip
	for _, rule := range wlists {
		services.RemoveIP(rule.Ip)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"message": "clear all rules",
	})
}

func AddIPRule(c *gin.Context) {
	domain := c.Param("domain")
	var ruleForm models.RuleForm
	if err := c.BindJSON(&ruleForm); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := services.AddIPRule(domain, &ruleForm); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// AllowIP
	services.AllowIP(ruleForm.Ip)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"message": "add rule success",
	})
}

func UpdateIPRule(c *gin.Context) {
	ruleid := c.Param("ruleid")
	ruleForm := models.RuleUpdateForm{}
	if err := c.BindJSON(&ruleForm); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := services.UpdateIPRule(ruleid, &ruleForm); err != nil {
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
	rid := c.Param("ruleid")
	rule := models.Whitelist{}

	if err := services.GetIpRule(rid, &rule); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := services.DeleteIPRule(rid); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// RemoveIP
	services.RemoveIP(rule.Ip)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"message": "delete rule success",
	})
}
