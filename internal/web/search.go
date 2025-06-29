package web

import (
	"log"

	"github.com/gin-gonic/gin"
)

var searchObj = gin.H{
	"Path": "/search",
}

func SearchNotes(ctx *gin.Context) {
	data := &gin.H{}

	query := ctx.DefaultQuery("q", "")

	notes, err := notesService.SearchNotes(&query)

	if err != nil {
		log.Println(err)
	}

	data = &gin.H{
		"Query": query,
		"Notes": notes,
	}

	MergeH(data, &searchObj)

	RenderHTML(ctx.Writer, "search", data)
}
