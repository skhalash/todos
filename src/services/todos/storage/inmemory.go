package storage

import "services/todos/model"

type inMemoryStorage struct {
	todos []model.Todo
}

func (s *inMemoryStorage) Add(t model.Todo) error {
	s.todos = append(s.todos, t)
	return nil
}

func (s inMemoryStorage) GetAll() ([]model.Todo, error) {
	return s.todos, nil
}

// New returns a new storage
func New() model.Storage {
	return &inMemoryStorage{}
}
