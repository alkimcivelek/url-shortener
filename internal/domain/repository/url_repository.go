package repository

import (
	"errors"
	"url-shortener/internal/domain/entity"
	"url-shortener/internal/domain/valueobject"
)

var (
	ErrURLNotFound  = errors.New("URL not found")
	ErrDuplicateURL = errors.New("URL already exists")
)

type URLRepository interface {
	Save(url *entity.URL) error
	FindByShortCode(shortCode valueobject.ShortCode) (*entity.URL, error)
	FindByOriginalURL(originalURL valueobject.OriginalURL) (*entity.URL, error)
	Update(url *entity.URL) error
	GetAll() ([]*entity.URL, error)
}
