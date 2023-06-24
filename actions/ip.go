package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uees/hidedomain/utils"
)

func GetIP(c *gin.Context) {
	ip := c.ClientIP()
	cidr := utils.GenIPV4Cidr(ip, 16)

	body := gin.H{
		"ip":      c.ClientIP(),
		"version": "IPv4",
		"network": cidr,
	}

	c.JSON(http.StatusOK, body)
}
