package service

import (
	"url-shortener/internal/application/port"
	"url-shortener/internal/application/usecase"
	"url-shortener/internal/domain/entity"
	"url-shortener/internal/domain/repository"
	domainService "url-shortener/internal/domain/service"
)

type URLApplicationService struct {
	shortenUseCase  *usecase.ShortenURLUseCase
	redirectUseCase *usecase.RedirectURLUseCase
	statsUseCase    *usecase.GetStatsUseCase
}

func NewURLApplicationService(repo repository.URLRepository) port.URLService {
	domainSvc := domainService.NewURLDomainService()

	return &URLApplicationService{
		shortenUseCase:  usecase.NewShortenURLUseCase(repo, domainSvc),
		redirectUseCase: usecase.NewRedirectURLUseCase(repo),
		statsUseCase:    usecase.NewGetStatsUseCase(repo),
	}
}

func (s *URLApplicationService) ShortenURL(originalURL string) (*entity.URL, error) {
	return s.shortenUseCase.Execute(originalURL)
}

func (s *URLApplicationService) RedirectURL(shortCode string) (*entity.URL, error) {
	return s.redirectUseCase.Execute(shortCode)
}

func (s *URLApplicationService) GetURLStats(shortCode string) (*entity.URL, error) {
	return s.statsUseCase.Execute(shortCode)
}
