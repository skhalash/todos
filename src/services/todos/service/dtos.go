package service

import "time"

// CreateTodoDto ...
type CreateTodoDto struct {
	Name        string    `json:"name"`
	Until       time.Time `json:"until"`
	Description string    `json:"decription"`
}
