package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/uees/hidedomain/actions"
	"github.com/uees/hidedomain/middlewares"
	"github.com/uees/hidedomain/utils"
)

func initRotuer() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Content-Type", "Access-Token", "Authorization", "X-Requested-With"}
	router.Use(cors.New(config)) // Default() allows all origins

	api := router.Group("/api")
	{
		api.POST("/token", actions.Login)
		api.GET("/ip", actions.GetIP)
		api.GET("/subscribe", actions.Subscription)

		authorized := api.Group("")
		authorized.Use(middlewares.AuthRequired())
		{
			authorized.GET("/profile", actions.Profile)

			authorized.GET("/domains", actions.DomainList)
			authorized.POST("/domains", actions.CreateDomain)
			authorized.GET("/domains/:domain", actions.ShowDomain)
			authorized.PATCH("/domains/:domain", actions.UpdateDomain)
			authorized.DELETE("domains/:domain", actions.DeleteDomain)

			authorized.GET("/domains/:domain/whitelist", actions.ShowList)
			authorized.DELETE("/domains/:domain/whitelist", actions.ClearList)
			authorized.POST("/domains/:domain/whitelist", actions.AddIPRule)

			authorized.GET("/domains/:domain/whitelist/:ruleid", actions.ShowIPRule)
			authorized.DELETE("/domains/:domain/whitelist/:ruleid", actions.DeleteIPRule)
			authorized.PATCH("/domains/:domain/whitelist/:ruleid", actions.UpdateIPRule)

			authorized.GET("/proxies", actions.ProxiesList)
			authorized.POST("/proxies", actions.CreateProxyitem)
			authorized.GET("/proxies/:id", actions.ShowProxyitem)
			authorized.PATCH("/proxies/:id", actions.UpdateProxyitem)
			authorized.DELETE("proxies/:id", actions.DeleteProxyItem)
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found!")
	})

	return router
}

func main() {
	utils.InitLoger()
	rotuer := initRotuer()
	rotuer.Run(utils.Conf.Listen)
}
