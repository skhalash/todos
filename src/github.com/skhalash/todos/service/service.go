package service

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	router.HandleFunc("/todos", getTodos).Methods("GET")

	log.Print("Launching the server")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
