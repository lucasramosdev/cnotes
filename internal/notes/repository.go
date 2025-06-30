package notes

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	RecentNotes(ctx context.Context) ([]BasicNote, error)
	GetNote(ctx context.Context, id *int64) (*Note, error)
	SearchNotes(ctx context.Context, search *string) ([]BasicNote, error)
	Create(ctx context.Context, note *Note) error
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

func (r *RepositoryPostgres) Create(ctx context.Context, note *Note) error {
	var category int64
	var theme int64
	if err := r.Conn.QueryRow(
		ctx,
		`SELECT id from categories where description = $1`,
		note.Category,
	).Scan(&category); err != nil {
		return fmt.Errorf("category %s not found", note.Category)
	}

	if err := r.Conn.QueryRow(
		ctx,
		`SELECT id from themes where description = $1`,
		note.Theme,
	).Scan(&theme); err != nil {
		return fmt.Errorf("theme %s not found", note.Theme)
	}

	_, err := r.Conn.Exec(
		ctx,
		`INSERT INTO notes (id, category, theme, title, summary) VALUES ($1, $2, $3, $4, $5)`,
		note.ID,
		category,
		theme,
		note.Title,
		note.Summary,
	)

	if err != nil {
		return err
	}

	for position, keyword := range note.Keywords {
		_, err := r.Conn.Exec(
			ctx,
			`INSERT INTO keywords (note, description, position) VALUES ($1, $2, $3)`,
			note.ID,
			keyword,
			position,
		)

		if err != nil {
			return err
		}
	}

	for position, annotation := range note.Annotations {
		_, err := r.Conn.Exec(
			ctx,
			`INSERT INTO annotations (note, value, position) VALUES ($1, $2, $3)`,
			note.ID,
			annotation,
			position,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
