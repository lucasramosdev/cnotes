package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", gin.H{})
}

func RedirectHome(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}
