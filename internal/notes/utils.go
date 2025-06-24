package notes

import "context"

func (r *RepositoryPostgres) getKeywords(ctx context.Context, id *int64, note *Note) error {
	keywords, err := r.Conn.Query(
		ctx,
		`SELECT description from keywords where note = $1
		order by position asc;`,
		*id,
	)

	if err != nil {
		return err
	}

	for keywords.Next() {
		var keyword string

		if err := keywords.Scan(&keyword); err != nil {
			return err
		}

		note.Keywords = append(note.Keywords, keyword)

	}

	return nil
}

func (r *RepositoryPostgres) getAnnotations(ctx context.Context, id *int64, note *Note) error {
	annotations, err := r.Conn.Query(
		ctx,
		`SELECT value from annotations where note = $1
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

		note.Annotations = append(note.Annotations, annotation)
	}

	return nil
}
