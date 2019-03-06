package model

import "errors"

// Name represents a todo name
type Name string

var (
	// ErrEmptyName signalizes that provided name was empty
	ErrEmptyName = errors.New("empty name")

	// ErrNameTooLong signalizes that provided name was too long
	ErrNameTooLong = errors.New("name too long")
)

var (
	maxNameLength = 100
)

// NewName parses a raw name
func NewName(raw string) (Name, error) {
	if len(raw) == 0 {
		return "", ErrEmptyName
	}

	if len(raw) > maxNameLength {
		return "", ErrNameTooLong
	}

	return "", nil
}
