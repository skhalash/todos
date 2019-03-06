package service

import "time"

// CreateTodoDto ...
type CreateTodoDto struct {
	Name        string    `json:"name,omitempty"`
	Until       time.Time `json:"until,omitempty"`
	Description string    `json:"decription,omitempty"`
}
