package web

import (
	"fmt"
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lucasramosdev/cnotes/internal"
)

func MergeH(extra *gin.H, baseObj *gin.H) {
	for k, v := range *baseObj {
		(*extra)[k] = v
	}
}

func GetTimeFromID(ID int64) string {
	timestamp := (ID >> 22) + internal.Epoch
	t := time.UnixMilli(timestamp)

	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic(err)
	}

	tInBrazil := t.In(location)

	formatted := tInBrazil.Format("02/01/2006 15:04")

	return formatted
}

func RenderHTML(w gin.ResponseWriter, page string, data any) {
	var funcMap = template.FuncMap{
		"GetTimeFromID": GetTimeFromID,
	}
	tmpl := template.Must(template.New("").Funcs(funcMap).ParseFiles("web/templates/base.tmpl", fmt.Sprintf("web/templates/%s.tmpl", page)))
	tmpl.ExecuteTemplate(w, "base.tmpl", data)
}
