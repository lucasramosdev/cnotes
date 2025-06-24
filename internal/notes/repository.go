package notes

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	RecentNotes(ctx context.Context) ([]BasicNote, error)
	GetNote(ctx context.Context, id *int64) (*Note, error)
	SearchNotes(ctx context.Context, search *string) ([]BasicNote, error)
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

	if err := r.getKeywords(ctx, id, &note); err != nil {
		return nil, err
	}

	if err := r.getAnnotations(ctx, id, &note); err != nil {
		return nil, err
	}

	return &note, nil

}

func (r *RepositoryPostgres) SearchNotes(ctx context.Context, search *string) ([]BasicNote, error) {
	var items []BasicNote

	rows, err := r.Conn.Query(
		ctx,
		`SELECT notes.id as id, notes.title as title, categories.description as category, themes.description as theme  from notes 
		left join categories on notes.category = categories.id
		left join themes on notes.theme = themes.id
		where 	concat_ws(' ', title, summary, category, theme) ILIKE '%' || $1 || '%'
		or exists (
			SELECT 1
			from keywords as kw
			where kw.description ILIKE '%' || $1 || '%'
		)
		or exists (
			SELECT 1
			from annotations as ant
			where ant.value ILIKE '%' || $1 || '%'
		)
		order by id desc;`,
		*search,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item BasicNote

		if err := rows.Scan(&item.ID, &item.Title, &item.Category, &item.Theme); err != nil {
			return nil, err
		}

		items = append(items, item)

	}

	return items, nil

}
