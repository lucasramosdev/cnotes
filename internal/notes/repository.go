package notes

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	RecentNotes(ctx context.Context) ([]BasicNote, error)
	GetNote(ctx context.Context, id *int64) (*Note, error)
}

type RepositoryPostgres struct {
	Conn *pgxpool.Pool
}

func (r *RepositoryPostgres) RecentNotes(ctx context.Context) ([]BasicNote, error) {
	rows, err := r.Conn.Query(
		ctx,
		`SELECT notes.id as id, notes.title as title, categories.description as category, themes.description as theme  from notes 
		left join categories on notes.category = categories.id
		left join themes on notes.theme = themes.id 
		order by id desc limit 20`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []BasicNote

	for rows.Next() {
		var item BasicNote

		if err := rows.Scan(&item.ID, &item.Title, &item.Category, &item.Theme); err != nil {
			return nil, err
		}

		items = append(items, item)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil

}

func (r *RepositoryPostgres) GetNote(ctx context.Context, id *int64) (*Note, error) {
	var note = Note{
		ID:          *id,
		Keywords:    []string{},
		Annotations: []string{},
	}

	err := r.Conn.QueryRow(
		ctx,
		`SELECT notes.title as title, categories.description as category, themes.description as theme, notes.summary as summary from notes 
		left join categories on notes.category = categories.id
		left join themes on notes.theme = themes.id
		where notes.id = $1;`,
		*id,
	).Scan(&note.Title, &note.Category, &note.Theme, &note.Summary)

	if err != nil {
		return nil, err
	}

	keywords, err := r.Conn.Query(
		ctx,
		`SELECT description from keywords where note = $1
		order by position asc;`,
		*id,
	)

	if err != nil {
		return nil, err
	}

	for keywords.Next() {
		var keyword string

		if err := keywords.Scan(&keyword); err != nil {
			return nil, err
		}

		note.Keywords = append(note.Keywords, keyword)

	}

	annotations, err := r.Conn.Query(
		ctx,
		`SELECT value from annotations where note = $1
		order by position asc;`,
		*id,
	)

	if err != nil {
		return nil, err
	}

	for annotations.Next() {
		var annotation string

		if err := annotations.Scan(&annotation); err != nil {
			return nil, err
		}

		note.Annotations = append(note.Annotations, annotation)
	}

	return &note, nil

}
