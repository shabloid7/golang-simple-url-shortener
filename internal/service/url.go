package service

import (
	"context"
	"fmt"
	"url-shortener/internal/errors"
	"url-shortener/internal/pkg/randstr"
	"url-shortener/internal/repository"
)

type URLService interface {
	Shorten(ctx context.Context, originalURL string) (string, error)
	Resolve(ctx context.Context, code string) (string, error)
}

type urlService struct {
	repo repository.URLRepository
	baseURL string
}

func NewURLService(repo repository.URLRepository, baseURL string) URLService {
    return &urlService{repo: repo, baseURL: baseURL}
}

func (s *urlService) Shorten(ctx context.Context, originalURL string) (string, error) {
	code := randstr.Generate(7)

	if err := s.repo.Save(ctx, code, originalURL); err != nil {
		return "", fmt.Errorf("%w: %w", errors.ErrURLSaveFailed, err)
	}
	
	return s.baseURL + "/" + code, nil
}

func (s *urlService) Resolve(ctx context.Context, code string) (string, error) {
	url, err := s.repo.Get(ctx, code)
	if err != nil {
		return "", fmt.Errorf("%w: %w", errors.ErrURLResolveFailed, err)
	}
	if url == "" {
		return "", errors.ErrURLNotFound
	}
	return url, nil
}


