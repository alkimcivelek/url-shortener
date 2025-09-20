package memory_test

import (
	"testing"
	"url-shortener/internal/domain/entity"
	"url-shortener/internal/domain/repository"
	"url-shortener/internal/domain/valueobject"
	"url-shortener/internal/infrastructure/persistence/memory"
)

func TestURLRepositoryMemory_Save(t *testing.T) {
	repo := memory.NewURLRepositoryMemory()

	url, err := entity.NewURL("https://test.com", "abc123")
	if err != nil {
		t.Fatalf("NewURL() failed: %v", err)
	}

	err = repo.Save(url)
	if err != nil {
		t.Errorf("Save() failed: %v", err)
	}

	url2, err := entity.NewURL("https://test.com", "def456")
	if err != nil {
		t.Fatalf("NewURL() failed: %v", err)
	}

	err = repo.Save(url2)
	if err != repository.ErrDuplicateURL {
		t.Errorf("Save() duplicate URL error = %v, want %v", err, repository.ErrDuplicateURL)
	}
}

func TestURLRepositoryMemory_FindByShortCode(t *testing.T) {
	repo := memory.NewURLRepositoryMemory()

	url, err := entity.NewURL("https://test.com", "test456")
	if err != nil {
		t.Fatalf("NewURL() failed: %v", err)
	}

	err = repo.Save(url)
	if err != nil {
		t.Fatalf("Save() failed: %v", err)
	}

	shortCode, _ := valueobject.NewShortCode("test456")
	found, err := repo.FindByShortCode(shortCode)
	if err != nil {
		t.Errorf("FindByShortCode() failed: %v", err)
	}
	if found.OriginalURL().Value() != url.OriginalURL().Value() {
		t.Errorf("FindByShortCode() URL = %v, want %v", found.OriginalURL().Value(), url.OriginalURL().Value())
	}

	nonExistentCode, _ := valueobject.NewShortCode("nonexistenturl")
	_, err = repo.FindByShortCode(nonExistentCode)
	if err != repository.ErrURLNotFound {
		t.Errorf("FindByShortCode() error = %v, want %v", err, repository.ErrURLNotFound)
	}
}

func TestURLRepositoryMemory_Update(t *testing.T) {
	repo := memory.NewURLRepositoryMemory()

	url, err := entity.NewURL("https://test.com", "test123")
	if err != nil {
		t.Fatalf("NewURL() failed: %v", err)
	}

	err = repo.Save(url)
	if err != nil {
		t.Fatalf("Save() failed: %v", err)
	}

	url.RecordAccess()

	err = repo.Update(url)
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	shortCode, _ := valueobject.NewShortCode("test123")
	found, err := repo.FindByShortCode(shortCode)
	if err != nil {
		t.Fatalf("FindByShortCode() failed: %v", err)
	}

	if found.AccessCount() != 1 {
		t.Errorf("Access count = %v, want 1", found.AccessCount())
	}
}
