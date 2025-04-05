package http

import (
	"github.com/lucasramosdev/cnotes/internal/database"
	"github.com/lucasramosdev/cnotes/internal/notes"
)

var notesService notes.Service

func Configure() {
	notesService = notes.Service{
		Repository: &notes.RepositoryPostgress{
			Conn: database.Conn,
		},
	}
}
