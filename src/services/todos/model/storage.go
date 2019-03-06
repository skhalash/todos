package model

type Storage interface {
	Add(t Todo) error

	GetAll() ([]Todo, error)
}
