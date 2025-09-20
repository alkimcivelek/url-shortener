package service

import (
	"url-shortener/internal/domain/entity"
	"url-shortener/pkg/hash"
)

type URLDomainService struct{}

func NewURLDomainService() *URLDomainService {
	return &URLDomainService{}
}

func (s *URLDomainService) GenerateShortCode(originalURL string) string {
	return hash.GenerateShortCode(originalURL)
}

func (s *URLDomainService) ValidateBusinessRules(url *entity.URL) error {
	// Domain-specific validation logic
	return nil
}
