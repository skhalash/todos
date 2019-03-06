package model

// Storage represents a todo storage
type Storage interface {
	Add(t Todo) error

	GetAll() ([]Todo, error)
}
