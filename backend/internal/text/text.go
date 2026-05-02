package text

import (
	"context"
	"fmt"

	"github.com/mgiks/typo-typer/internal/storage"
)

type TextService struct {
	repo storage.TextRepository
}

func NewService(repo storage.TextRepository) TextService {
	return TextService{
		repo: repo,
	}
}

func (s TextService) GetRandomText(ctx context.Context) (string, error) {
	text, err := s.repo.GetRandomText(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get random text from repo: %w", err)
	}
	return text, nil
}
