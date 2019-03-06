package model

import (
	"errors"
	"time"
)

// Name represents a todo name
type Name string

// Description represents a todo description
type Description string

// Todo represents a single todo item
type Todo struct {
	Name        Name
	Description Description
	CreatedAt   time.Time
	Until       time.Time
}

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

// NewTodo ...
func NewTodo(name, description string, createdAt, until time.Time) (*Todo, error) {
	n, err := newName(name)
	if err != nil {
		return nil, err
	}

	d, err := newDescription(description)
	if err != nil {
		return nil, err
	}

	return &Todo{
		Name:        n,
		Description: d,
		CreatedAt:   createdAt,
		Until:       until,
	}, nil
}

func newName(raw string) (Name, error) {
	if len(raw) == 0 {
		return "", ErrEmptyName
	}

	if len(raw) > maxNameLength {
		return "", ErrNameTooLong
	}

	return Name(raw), nil
}

func newDescription(raw string) (Description, error) {
	if len(raw) > maxDecriptionLength {
		return "", ErrDescriptionTooLong
	}

	return Description(raw), nil
}
