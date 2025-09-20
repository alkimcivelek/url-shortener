package usecase_test

import (
	"testing"
	"url-shortener/internal/application/usecase"
	"url-shortener/internal/domain/service"
	"url-shortener/internal/infrastructure/persistence/memory"
)

func TestShortenURLUseCase_Execute(t *testing.T) {
	repo := memory.NewURLRepositoryMemory()
	domainService := service.NewURLDomainService()
	useCase := usecase.NewShortenURLUseCase(repo, domainService)

	tests := []struct {
		name        string
		originalURL string
		wantErr     bool
	}{
		{
			name:        "Valid URL",
			originalURL: "https://test.com/path",
			wantErr:     false,
		},
		{
			name:        "Invalid URL",
			originalURL: "invalid-url",
			wantErr:     true,
		},
		{
			name:        "Empty URL",
			originalURL: "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := useCase.Execute(tt.originalURL)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Execute() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("Execute() unexpected error = %v", err)
				return
			}

			if result == nil {
				t.Errorf("Execute() result is nil")
				return
			}

			if result.OriginalURL().Value() != tt.originalURL {
				t.Errorf("Execute() original URL = %v, want %v", result.OriginalURL().Value(), tt.originalURL)
			}

			if result.ShortCode().Value() == "" {
				t.Errorf("Execute() short code is empty")
			}
		})
	}
}

func TestShortenURLUseCase_ExecuteDuplicate(t *testing.T) {
	repo := memory.NewURLRepositoryMemory()
	domainService := service.NewURLDomainService()
	useCase := usecase.NewShortenURLUseCase(repo, domainService)

	originalURL := "https://test.com"

	result1, err := useCase.Execute(originalURL)
	if err != nil {
		t.Fatalf("First Execute() failed: %v", err)
	}

	result2, err := useCase.Execute(originalURL)
	if err != nil {
		t.Fatalf("Second Execute() failed: %v", err)
	}

	if result1.ShortCode().Value() != result2.ShortCode().Value() {
		t.Errorf("Duplicate URL should return same short code: %v != %v",
			result1.ShortCode().Value(), result2.ShortCode().Value())
	}
}
