package notes

import "context"

type Service struct {
	Repository Repository
}

func (s Service) RecentNotes(ctx context.Context) ([]BasicNote, error) {
	notes, err := s.Repository.RecentNotes(ctx)
	if err != nil {
		return nil, err
	}

	return notes, nil
}
