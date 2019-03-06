package service

import (
	"log"
	"net/http"
)

// Service represents a todo web service
type Service struct{}

// NewService creates a new instance of Service
func NewService() *Service {
	return &Service{}
}

// Run launches the Service
func (*Service) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/todos", handleGetTodos).Methods("GET")

	log.Print("Launching the server")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func handleGetTodos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
