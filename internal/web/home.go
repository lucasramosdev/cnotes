package web

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lucasramosdev/cnotes/internal/notes"
)

var homeObj = &gin.H{
	"Path": "/home",
}

func GetHome(ctx *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	items, err := notesService.RecentNotes(ctxTimeout)
	if err != nil {
		log.Println(err.Error())
		items = []notes.BasicNote{}
	}
	data := &gin.H{
		"Notes": items,
	}
	MergeH(data, homeObj)

	RenderHTML(ctx.Writer, "home", data)
}

func RedirectHome(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "/")
}
