package model

import "errors"

// Name represents a todo name
type Name string

// Description represents a todo description
type Description string

var (
	// ErrEmptyName signalizes that provided name was empty
	ErrEmptyName = errors.New("empty name")

	// ErrNameTooLong signalizes that provided name was too long
	ErrNameTooLong = errors.New("name too long")

	// ErrDescriptionTooLong signalizes that provided description was too long
	ErrDescriptionTooLong = errors.New("description too long")
)

var (
	maxNameLength       = 100
	maxDecriptionLength = 300
)

// NewName parses a raw name
func NewName(raw string) (Name, error) {
	if len(raw) == 0 {
		return "", ErrEmptyName
	}

	if len(raw) > maxNameLength {
		return "", ErrNameTooLong
	}

	return Name(raw), nil
}

// NewDescription parses a raw description
func NewDescription(raw string) (Description, error) {
	if len(raw) > maxDecriptionLength {
		return "", ErrDescriptionTooLong
	}

	return Description(raw), nil
}
