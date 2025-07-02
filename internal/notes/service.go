package notes

import (
	"context"
	"html/template"
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

func (s Service) GetNote(ctx context.Context, id *int64) (*NoteDetails, error) {
	note, err := s.Repository.GetNote(ctx, id)
	if err != nil {
		return nil, err
	}

	processedClues := make([]struct {
		Value       string
		Annotations []template.HTML
	}, 0, len(note.Clues))

	for _, clue := range note.Clues {
		var annotations []template.HTML
		for _, annotation := range clue.Annotations {
			annotations = append(annotations, template.HTML(annotation))
		}
		processedClues = append(processedClues, struct {
			Value       string
			Annotations []template.HTML
		}{Value: clue.Value, Annotations: annotations})
	}
	return &NoteDetails{
		ID:       note.ID,
		Category: note.Category,
		Theme:    note.Theme,
		Title:    note.Title,
		Summary:  note.Summary,
		Clues:    processedClues,
	}, nil
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
