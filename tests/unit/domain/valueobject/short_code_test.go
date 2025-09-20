package valueobject_test

import (
	"testing"
	"url-shortener/internal/domain/valueobject"
)

func TestNewShortCode(t *testing.T) {
	tests := []struct {
		name    string
		code    string
		wantErr bool
		errType error
	}{
		{"Valid short code", "abc123", false, nil},
		{"Valid alphanumeric", "xyz789", false, nil},
		{"Empty code", "", true, valueobject.ErrEmptyShortCode},
		{"Whitespace code", "   ", true, valueobject.ErrEmptyShortCode},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shortCode, err := valueobject.NewShortCode(tt.code)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewShortCode() expected error, got nil")
					return
				}
				if tt.errType != nil && err != tt.errType {
					t.Errorf("NewShortCode() error = %v, want %v", err, tt.errType)
				}
				return
			}

			if err != nil {
				t.Errorf("NewShortCode() unexpected error = %v", err)
				return
			}

			if shortCode.Value() != tt.code {
				t.Errorf("NewShortCode() value = %v, want %v", shortCode.Value(), tt.code)
			}
		})
	}
}
