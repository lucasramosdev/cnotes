package http

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lucasramosdev/cnotes/internal/notes"
)

var baseObj = gin.H{
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
	mergeH(data)
	ctx.HTML(http.StatusOK, "home.tmpl", data)
}

func RedirectHome(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}
