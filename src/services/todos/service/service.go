package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"services/todos/model"
	"time"

	"github.com/gorilla/mux"
)

// Service represents a todo web service
type Service struct {
	storage model.Storage
}

// NewService creates a new instance of Service
func NewService(s model.Storage) *Service {
	return &Service{storage: s}
}

// Run launches the Service
func (s Service) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/todos", s.handleGetTodos).Methods("GET")
	router.HandleFunc("/todos", s.handleCreateTodo).Methods("POST")

	log.Print("Launching the server")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func (s Service) handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var request CreateTodoRequest
	if err = json.Unmarshal(body, &request); err != nil {
		http.Error(w, "Error unmarshaling json", http.StatusBadRequest)
		return
	}

	todo, err := model.NewTodo(request.Name, request.Description, time.Now().UTC(), request.Until)
	if err != nil {
		if err == model.ErrEmptyName {
			http.Error(w, "Empty name", http.StatusBadRequest)
			return
		}

		if err == model.ErrNameTooLong {
			http.Error(w, "Name too long", http.StatusBadRequest)
			return
		}

		if err == model.ErrDescriptionTooLong {
			http.Error(w, "Description too long", http.StatusBadRequest)
			return
		}
	}

	if err = s.storage.Add(*todo); err != nil {
		http.Error(w, "Error accessing the storage", http.StatusInternalServerError)
		return
	}
}

func (s Service) handleGetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := s.storage.GetAll()
	if err != nil {
		http.Error(w, "Error accessing the storage", http.StatusInternalServerError)
		return
	}

	response := GetTodosResponse{
		Todos: make([]Todo, len(todos), len(todos)),
	}
	for i, todo := range todos {
		response.Todos[i] = toTodoDto(todo)
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error marshaling json", http.StatusInternalServerError)
	}

	w.Write(bytes)
}

func toTodoDto(todo model.Todo) Todo {
	return Todo{
		Name:        string(todo.Name),
		Description: string(todo.Description),
		CreatedAt:   todo.CreatedAt,
		Until:       todo.Until,
	}
}
