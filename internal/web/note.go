package web

import (
	"log"

	"github.com/gin-gonic/gin"
)

var noteObj = gin.H{
	"Path": "/notes",
}

func GetNote(ctx *gin.Context) {
	// ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*1)
	// defer cancel()

	id := ctx.Param("id")
	log.Println(id)

	RenderHTML(ctx.Writer, "note", &gin.H{"ID": id})
}
