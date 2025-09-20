package valueobject

import (
	"errors"
	"net/url"
	"strings"
)

var (
	ErrEmptyURL          = errors.New("URL cannot be empty")
	ErrInvalidURL        = errors.New("invalid URL format")
	ErrUnsupportedScheme = errors.New("unsupported URL scheme")
)

type OriginalURL struct {
	value string
}

func NewOriginalURL(rawURL string) (OriginalURL, error) {
	if strings.TrimSpace(rawURL) == "" {
		return OriginalURL{}, ErrEmptyURL
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return OriginalURL{}, ErrInvalidURL
	}

	if parsedURL.Scheme == "" {
		return OriginalURL{}, ErrInvalidURL
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return OriginalURL{}, ErrUnsupportedScheme
	}

	if parsedURL.Host == "" {
		return OriginalURL{}, ErrInvalidURL
	}

	return OriginalURL{value: rawURL}, nil
}

func (ou OriginalURL) Value() string {
	return ou.value
}
