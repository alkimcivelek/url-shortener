package usecase

import (
	"url-shortener/internal/domain/entity"
	"url-shortener/internal/domain/repository"
	"url-shortener/internal/domain/service"
	"url-shortener/internal/domain/valueobject"
)

type ShortenURLUseCase struct {
	urlRepository repository.URLRepository
	domainService *service.URLDomainService
}

func NewShortenURLUseCase(repo repository.URLRepository, domainService *service.URLDomainService) *ShortenURLUseCase {
	return &ShortenURLUseCase{
		urlRepository: repo,
		domainService: domainService,
	}
}

func (uc *ShortenURLUseCase) Execute(originalURLStr string) (*entity.URL, error) {
	originalURL, err := valueobject.NewOriginalURL(originalURLStr)
	if err != nil {
		return nil, err
	}

	existingURL, err := uc.urlRepository.FindByOriginalURL(originalURL)
	if err == nil {
		return existingURL, nil
	}

	shortCodeStr := uc.domainService.GenerateShortCode(originalURLStr)

	url, err := entity.NewURL(originalURLStr, shortCodeStr)
	if err != nil {
		return nil, err
	}

	if err := uc.domainService.ValidateBusinessRules(url); err != nil {
		return nil, err
	}

	if err := uc.urlRepository.Save(url); err != nil {
		return nil, err
	}

	return url, nil
}
