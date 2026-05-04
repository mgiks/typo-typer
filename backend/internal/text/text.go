package text

import (
	"context"
	"fmt"

	"github.com/mgiks/typo-typer/internal/storage"
)

type TextService interface {
	GetRandomText(context.Context) (storage.Text, error)
}

type textService struct {
	repo storage.TextRepository
}

func NewService(repo storage.TextRepository) TextService {
	return textService{
		repo: repo,
	}
}

func (s textService) GetRandomText(ctx context.Context) (storage.Text, error) {
	text, err := s.repo.GetRandomText(ctx)
	if err != nil {
		return storage.Text{}, fmt.Errorf("failed to get random text from repo: %w", err)
	}
	return text, nil
}
