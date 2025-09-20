package valueobject_test

import (
	"testing"
	"url-shortener/internal/domain/valueobject"
)

func TestNewOriginalURL(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
		errType error
	}{
		{"Valid HTTP URL", "http://test.com", false, nil},
		{"Valid HTTPS URL", "https://test.com", false, nil},
		{"Valid URL with path", "https://test.com/path", false, nil},
		{"Valid URL with query", "https://test.com?query=1", false, nil},
		{"Empty URL", "", true, valueobject.ErrEmptyURL},
		{"Whitespace URL", "   ", true, valueobject.ErrEmptyURL},
		{"Invalid scheme", "ftp://test.com", true, valueobject.ErrUnsupportedScheme},
		{"No scheme", "test.com", true, valueobject.ErrInvalidURL},
		{"No host", "http://", true, valueobject.ErrInvalidURL},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalURL, err := valueobject.NewOriginalURL(tt.url)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewOriginalURL() expected error, got nil")
					return
				}
				if tt.errType != nil && err != tt.errType {
					t.Errorf("NewOriginalURL() error = %v, want %v", err, tt.errType)
				}
				return
			}

			if err != nil {
				t.Errorf("NewOriginalURL() unexpected error = %v", err)
				return
			}

			if originalURL.Value() != tt.url {
				t.Errorf("NewOriginalURL() value = %v, want %v", originalURL.Value(), tt.url)
			}
		})
	}
}
