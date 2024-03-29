package actions

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uees/hidedomain/models"
	"github.com/uees/hidedomain/services"
	"gorm.io/gorm"
)

func DomainList(c *gin.Context) {
	var domains []models.Domain

	if err := services.GetAllDomains(&domains); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"data":    domains,
	})
}

func ShowDomain(c *gin.Context) {
	domainName := c.Param("domain")
	domain := models.Domain{}
	if err := services.QueryDomainByName(domainName, &domain); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"data":    domain,
	})
}

func CreateDomain(c *gin.Context) {
	postData := models.DomainForm{}
	if err := c.BindJSON(&postData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if ok, _ := services.HasDomain(postData.Name); ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"status":  "failed",
			"message": "domain already exists.",
		})
		return
	}

	if err := services.CreateDomain(&postData); err != nil {
		log.Fatalln("Create Domain err")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "success",
		"message": "add domain success.",
	})
}

func UpdateDomain(c *gin.Context) {
	domainName := c.Param("domain")

	var formData models.DomainForm
	c.BindJSON(&formData)

	if err := services.UpdateDomainByName(domainName, &formData); err != nil {
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
		"message": "update domain success.",
	})
}

func DeleteDomain(c *gin.Context) {
	domainName := c.Param("domain")

	if err := services.DeleteDomainByName(domainName); err != nil {
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
		"message": "delete domain success.",
	})
}
