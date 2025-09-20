package entity

import (
	"errors"
	"time"
	"url-shortener/internal/domain/valueobject"
)

var (
	ErrInvalidURL = errors.New("invalid URL")
	ErrEmptyURL   = errors.New("URL cannot be empty")
)

type URL struct {
	id          string
	originalURL valueobject.OriginalURL
	shortCode   valueobject.ShortCode
	createdAt   time.Time
	accessCount int64
}

func (u *URL) ID() string                           { return u.id }
func (u *URL) OriginalURL() valueobject.OriginalURL { return u.originalURL }
func (u *URL) ShortCode() valueobject.ShortCode     { return u.shortCode }
func (u *URL) CreatedAt() time.Time                 { return u.createdAt }
func (u *URL) AccessCount() int64                   { return u.accessCount }

func NewURL(originalURLStr, shortCodeStr string) (*URL, error) {
	originalURL, err := valueobject.NewOriginalURL(originalURLStr)
	if err != nil {
		return nil, err
	}

	shortCode, err := valueobject.NewShortCode(shortCodeStr)
	if err != nil {
		return nil, err
	}

	return &URL{
		id:          generateID(),
		originalURL: originalURL,
		shortCode:   shortCode,
		createdAt:   time.Now(),
		accessCount: 0,
	}, nil
}

func (u *URL) RecordAccess() {
	u.accessCount++
}

func generateID() string {
	return time.Now().Format("20060102150405") + randomString(6)
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(result)
}
