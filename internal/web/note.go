package web

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var noteObj = gin.H{
	"Path": "/notes",
}

func GetNote(ctx *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	data := &gin.H{}

	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Println(err)
	}

	note, err := notesService.GetNote(ctxTimeout, &id)

	if err != nil {
		log.Println(err)
	}

	data = &gin.H{
		"ID":          id,
		"Title":       note.Title,
		"Summary":     note.Summary,
		"Annotations": note.Annotations,
		"Keywords":    note.Keywords,
	}

	MergeH(data, &noteObj)

	RenderHTML(ctx.Writer, "note", data)
}
