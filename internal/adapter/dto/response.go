package dto

import (
	"time"
	"url-shortener/internal/domain/entity"
)

type ShortenURLResponse struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:"short_code"`
	CreatedAt   string `json:"created_at"`
}

type StatsResponse struct {
	ShortCode   string `json:"short_code"`
	OriginalURL string `json:"original_url"`
	AccessCount int64  `json:"access_count"`
	CreatedAt   string `json:"created_at"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type HealthResponse struct {
	Status    string `json:"status"`
	Service   string `json:"service"`
	Timestamp string `json:"timestamp"`
}

func URLEntityToShortenResponse(url *entity.URL, baseURL string) ShortenURLResponse {
	return ShortenURLResponse{
		ShortURL:    baseURL + "/" + url.ShortCode().Value(),
		OriginalURL: url.OriginalURL().Value(),
		ShortCode:   url.ShortCode().Value(),
		CreatedAt:   url.CreatedAt().Format(time.RFC3339),
	}
}

func URLEntityToStatsResponse(url *entity.URL) StatsResponse {
	return StatsResponse{
		ShortCode:   url.ShortCode().Value(),
		OriginalURL: url.OriginalURL().Value(),
		AccessCount: url.AccessCount(),
		CreatedAt:   url.CreatedAt().Format(time.RFC3339),
	}
}
