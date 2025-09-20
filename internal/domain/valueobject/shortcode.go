package valueobject

import (
	"errors"
	"strings"
)

var ErrEmptyShortCode = errors.New("short code cannot be empty")

type ShortCode struct {
	value string
}

func NewShortCode(code string) (ShortCode, error) {
	if strings.TrimSpace(code) == "" {
		return ShortCode{}, ErrEmptyShortCode
	}

	return ShortCode{value: code}, nil
}

func (sc ShortCode) Value() string {
	return sc.value
}
