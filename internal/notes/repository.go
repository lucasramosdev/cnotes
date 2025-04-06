package notes

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	RecentNotes(ctx context.Context) ([]BasicNote, error)
}

type RepositoryPostgress struct {
	Conn *pgxpool.Pool
}

func (r *RepositoryPostgress) RecentNotes(ctx context.Context) ([]BasicNote, error) {
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
