package usecase

import (
	"url-shortener/internal/domain/entity"
	"url-shortener/internal/domain/repository"
	"url-shortener/internal/domain/valueobject"
)

type GetStatsUseCase struct {
	urlRepository repository.URLRepository
}

func NewGetStatsUseCase(repo repository.URLRepository) *GetStatsUseCase {
	return &GetStatsUseCase{urlRepository: repo}
}

func (uc *GetStatsUseCase) Execute(shortCodeStr string) (*entity.URL, error) {
	shortCode, err := valueobject.NewShortCode(shortCodeStr)
	if err != nil {
		return nil, err
	}

	return uc.urlRepository.FindByShortCode(shortCode)
}
