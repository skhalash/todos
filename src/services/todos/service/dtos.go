package service

import "time"

// CreateTodoRequest ...
type CreateTodoRequest struct {
	Name        string    `json:"name,omitempty"`
	Until       time.Time `json:"until,omitempty"`
	Description string    `json:"decription,omitempty"`
}

// GetTodosResponse ...
type GetTodosResponse struct {
	Todos []Todo `json:"todos"`
}

// Todo ...
type Todo struct {
	Name        string    `json:"name,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	Until       time.Time `json:"until,omitempty"`
	Description string    `json:"decription,omitempty"`
}
