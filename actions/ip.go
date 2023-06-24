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
		"IP":      c.ClientIP(),
		"Version": "IPv4",
		"Network": cidr,
	}

	c.JSON(http.StatusOK, body)
}
