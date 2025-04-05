package http

import "github.com/gin-gonic/gin"

func mergeH(extra *gin.H) {
	for k, v := range baseObj {
		(*extra)[k] = v
	}
}
