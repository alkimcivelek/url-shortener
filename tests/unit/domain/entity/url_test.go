package entity_test

import (
	"testing"
	"url-shortener/internal/domain/entity"
)

func TestNewURL(t *testing.T) {
	tests := []struct {
		name        string
		originalURL string
		shortCode   string
		wantErr     bool
	}{
		{
			name:        "Valid URL and code",
			originalURL: "https://test.com",
			shortCode:   "abc123",
			wantErr:     false,
		},
		{
			name:        "Invalid URL",
			originalURL: "not-a-url",
			shortCode:   "abc123",
			wantErr:     true,
		},
		{
			name:        "Empty short code",
			originalURL: "https://test.com",
			shortCode:   "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url, err := entity.NewURL(tt.originalURL, tt.shortCode)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewURL() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("NewURL() unexpected error = %v", err)
				return
			}

			if url.OriginalURL().Value() != tt.originalURL {
				t.Errorf("NewURL() original URL = %v, want %v", url.OriginalURL().Value(), tt.originalURL)
			}

			if url.ShortCode().Value() != tt.shortCode {
				t.Errorf("NewURL() short code = %v, want %v", url.ShortCode().Value(), tt.shortCode)
			}

			if url.AccessCount() != 0 {
				t.Errorf("NewURL() access count = %v, want 0", url.AccessCount())
			}
		})
	}
}

func TestURL_RecordAccess(t *testing.T) {
	url, err := entity.NewURL("https://test.com", "abc123")
	if err != nil {
		t.Fatalf("NewURL() failed: %v", err)
	}

	if url.AccessCount() != 0 {
		t.Errorf("Initial access count = %v, want 0", url.AccessCount())
	}

	for i := 1; i <= 5; i++ {
		url.RecordAccess()
		if url.AccessCount() != int64(i) {
			t.Errorf("Access count after %d records = %v, want %v", i, url.AccessCount(), i)
		}
	}
}
