package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uees/hidedomain/models"
	"github.com/uees/hidedomain/services"
)

func ShowList(c *gin.Context) {
	domain := c.Param("domain")
	wlists := []models.Rule{}
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
	var ruleForm models.RuleForm
	if err := c.BindJSON(&ruleForm); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := services.AddIPRule(domain, &ruleForm); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"message": "add rule success",
	})
}

func ShowIPRule(c *gin.Context) {
	rid := c.Param("ruleid")
	domain := c.Param("domain")
	ipRule := models.Rule{}
	if err := services.GetIpRule(domain, rid, &ipRule); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"data":    ipRule,
	})
}

func UpdateIPRule(c *gin.Context) {
	ruleid := c.Param("ruleid")
	domain := c.Param("domain")
	ruleForm := models.RuleUpdateForm{}
	if err := c.BindJSON(&ruleForm); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := services.UpdateIPRule(domain, ruleid, &ruleForm); err != nil {
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
	domain := c.Param("domain")

	if err := services.DeleteIPRule(domain, rid); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"message": "delete rule success",
	})
}
