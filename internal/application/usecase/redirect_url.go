package usecase

import (
	"url-shortener/internal/domain/entity"
	"url-shortener/internal/domain/repository"
	"url-shortener/internal/domain/valueobject"
)

type RedirectURLUseCase struct {
	urlRepository repository.URLRepository
}

func NewRedirectURLUseCase(repo repository.URLRepository) *RedirectURLUseCase {
	return &RedirectURLUseCase{urlRepository: repo}
}

func (uc *RedirectURLUseCase) Execute(shortCodeStr string) (*entity.URL, error) {
	shortCode, err := valueobject.NewShortCode(shortCodeStr)
	if err != nil {
		return nil, err
	}

	url, err := uc.urlRepository.FindByShortCode(shortCode)
	if err != nil {
		return nil, err
	}

	url.RecordAccess()

	if err := uc.urlRepository.Update(url); err != nil {
		return nil, err
	}

	return url, nil
}
