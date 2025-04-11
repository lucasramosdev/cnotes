package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine) {
	g.GET("/", GetHome)
	g.GET("/home", RedirectHome)

	g.GET("/notes/:id", GetNote)

	g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
