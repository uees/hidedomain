package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/uees/hidedomain/actions"
	"github.com/uees/hidedomain/middlewares"
	"github.com/uees/hidedomain/services"
)

func initRotuer() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default()) // Default() allows all origins

	api := router.Group("/api")
	{
		api.POST("/token", actions.Login)

		authorized := api.Group("")
		authorized.Use(middlewares.AuthRequired())
		{
			authorized.GET("/profile", actions.Profile)

			authorized.GET("/domains", actions.DomainList)
			authorized.POST("/domains", actions.CreateDomain)
			authorized.PATCH("/domains/:domain", actions.UpdateDomain)
			authorized.DELETE("domains/:domain", actions.DeleteDomain)

			authorized.GET("/domains/:domain/whitelist", actions.ShowList)
			authorized.DELETE("/domains/:domain/whitelist", actions.ClearList)
			authorized.POST("/domains/:domain/whitelist", actions.AddIPRule)
			authorized.DELETE("/domains/:domain/whitelist/:ruleid", actions.DeleteIPRule)
			authorized.PATCH("/domains/:domain/whitelist/:ruleid", actions.UpdateIPRule)
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	return router
}

func main() {
	rotuer := initRotuer()
	rotuer.Run(services.Conf.Listen)
}
