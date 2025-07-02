package web

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lucasramosdev/cnotes/internal/notes"
)

var noteObj = gin.H{
	"Path": "/notes",
}

func GetNote(ctx *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Println(err)
	}

	note, err := notesService.GetNote(ctxTimeout, &id)

	if err != nil {
		log.Println(err)
	}

	data := &gin.H{
		"ID":       id,
		"Title":    note.Title,
		"Summary":  note.Summary,
		"Clues":    note.Clues,
		"Category": note.Category,
		"Theme":    note.Theme,
	}

	MergeH(data, &noteObj)

	RenderHTML(ctx.Writer, "note", data)
}

func CreateNote(ctx *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	var note notes.CreateNote

	if err := ctx.ShouldBindJSON(&note); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := notesService.Create(ctxTimeout, &note)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"id": id,
	},
	)

}
