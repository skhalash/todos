package main

// Service represents a todo web service
type Service struct{}

// NewService creates a new instance of Service
func NewService() *Service {
	return &Service{}
}

// Run launches the Service
func (*Service) Run() {

}
