package notes

import (
	"context"
	"time"
)

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

func (s Service) GetNote(ctx context.Context, id *int64) (*Note, error) {
	return s.Repository.GetNote(ctx, id)
}

func (s Service) SearchNotes(search *string) ([]BasicNote, error) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	if *search == "" {
		return s.RecentNotes(ctxTimeout)
	}

	return s.Repository.SearchNotes(ctxTimeout, search)
}
