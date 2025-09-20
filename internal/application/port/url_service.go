package port

import "url-shortener/internal/domain/entity"

type URLService interface {
	ShortenURL(originalURL string) (*entity.URL, error)
	RedirectURL(shortCode string) (*entity.URL, error)
	GetURLStats(shortCode string) (*entity.URL, error)
}
