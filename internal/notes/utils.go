package notes

import "context"

type QueryClue struct {
	ID    uint32 `json:"id"`
	Value string `json:"value"`
}

func (r *RepositoryPostgres) getClues(ctx context.Context, id *int64, note *Note) error {
	clues, err := r.Conn.Query(
		ctx,
		`SELECT id, value from clues where note = $1
		order by position asc;`,
		*id,
	)

	if err != nil {
		return err
	}

	for clues.Next() {
		var clue Clue
		var id uint32

		if err := clues.Scan(&id, &clue.Value); err != nil {
			return err
		}

		err := r.getAnnotations(ctx, &id, &clue)
		if err != nil {
			return err
		}

		note.Clues = append(note.Clues, clue)

	}

	return nil
}

func (r *RepositoryPostgres) getAnnotations(ctx context.Context, id *uint32, clue *Clue) error {
	annotations, err := r.Conn.Query(
		ctx,
		`SELECT value from annotations where clue = $1
		order by position asc;`,
		*id,
	)

	if err != nil {
		return err
	}

	for annotations.Next() {
		var annotation string

		if err := annotations.Scan(&annotation); err != nil {
			return err
		}

		clue.Annotations = append(clue.Annotations, annotation)
	}

	return nil
}
