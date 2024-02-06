package actions

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/uees/hidedomain/models"
	"github.com/uees/hidedomain/services"
	"gorm.io/gorm"
)

func ProxiesList(c *gin.Context) {
	var proxies = &[]models.Proxyitem{}

	if err := services.GetAllProxies(proxies); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"data":    proxies,
	})
}

func ShowProxyitem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var proxyItem = &models.Proxyitem{}
	if err := services.QueryProxyItem(uint(id), proxyItem); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"data":    proxyItem,
	})
}

func CreateProxyitem(c *gin.Context) {
	postData := models.ProxyitemForm{}
	if err := c.BindJSON(&postData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := services.CreateProxyItem(&postData); err != nil {
		log.Fatalln("Create ProxyItem err")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"message": "add proxyItem success.",
	})
}

func UpdateProxyitem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var formData models.ProxyitemForm
	c.BindJSON(&formData)

	if err := services.UpdateProxyItem(uint(id), &formData); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"message": "update proxyItem success.",
	})
}

func DeleteProxyItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := services.DeleteProxyItem(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"message": "delete proxyItem success.",
	})
}
