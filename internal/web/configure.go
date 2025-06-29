package web

import (
	"github.com/lucasramosdev/cnotes/internal/database"
	"github.com/lucasramosdev/cnotes/internal/notes"
	"github.com/lucasramosdev/cnotes/internal/users"
)

var notesService notes.Service
var usersService users.Service

func Configure() {
	notesService = notes.Service{
		Repository: &notes.RepositoryPostgres{
			Conn: database.Conn,
		},
	}

	usersService = users.Service{
		Repository: &users.RepositoryPostgres{
			Conn: database.Conn,
		},
	}
}
