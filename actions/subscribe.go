package actions

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/uees/hidedomain/models"
	"github.com/uees/hidedomain/services"
	"github.com/uees/hidedomain/utils"
)

// Subscription 输出订阅的内容
func Subscription(c *gin.Context) {

	token, ok := c.GetQuery("token")
	if !ok || token != utils.Conf.SubscribeToken {
		c.String(http.StatusForbidden, "Forbidden")
		return
	}

	var proxies = &[]models.Proxyitem{}
	err := services.GetAllResolvedProxies(proxies)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resultSlice := []string{}
	for _, proxy := range *proxies {
		proxyString := fmt.Sprintf("%s://%s", proxy.Protocol, base64.StdEncoding.EncodeToString([]byte(proxy.Content)))
		resultSlice = append(resultSlice, proxyString)
	}

	resultString := strings.Join(resultSlice, "\n")

	c.String(http.StatusOK, resultString)
}
