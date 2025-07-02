package notes

import (
	"context"
	"time"

	"github.com/lucasramosdev/cnotes/internal"
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
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	if *search == "" {
		return s.RecentNotes(ctxTimeout)
	}

	return s.Repository.SearchNotes(ctxTimeout, search)
}

func (s Service) Create(ctx context.Context, data *CreateNote) (*internal.ID, error) {
	node := internal.NewSnowflakeNode(1)
	id := node.GenerateID()
	note := &Note{
		ID:       int64(id),
		Category: data.Category,
		Theme:    data.Theme,
		Title:    data.Title,
		Summary:  data.Summary,
		Clues:    data.Clues,
	}

	if err := s.Repository.Create(ctx, note); err != nil {
		return nil, err
	}

	return &id, nil

}
