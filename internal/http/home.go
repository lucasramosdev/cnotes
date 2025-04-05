package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var baseObj = gin.H{
	"path": "/home",
}

func mergeH(extra *gin.H) {
	for k, v := range baseObj {
		(*extra)[k] = v
	}
}

func GetHome(c *gin.Context) {
	data := &gin.H{}
	mergeH(data)
	c.HTML(http.StatusOK, "home.tmpl", data)
}

func RedirectHome(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}
