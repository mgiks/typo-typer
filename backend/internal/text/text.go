package text

import (
	"context"

	"github.com/mgiks/typo-typer/internal/storage"
)

type TextService interface {
	GetRandomText(context.Context) (storage.Text, error)
	CreateText(context.Context, *storage.Text) error
}

type textService struct {
	text storage.TextRepository
}

func NewService(repo storage.TextRepository) TextService {
	return textService{
		text: repo,
	}
}

func (s textService) GetRandomText(ctx context.Context) (storage.Text, error) {
	text, err := s.text.GetRandom(ctx)
	if err != nil {
		return storage.Text{}, err
	}
	return text, nil
}

func (s textService) CreateText(ctx context.Context, text *storage.Text) error {
	if err := s.text.Create(ctx, text); err != nil {
		return err
	}
	return nil
}
