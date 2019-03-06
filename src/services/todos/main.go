package main

import (
	"services/todos/service"
	"services/todos/storage"
)

func main() {
	storage := storage.New()
	service := service.NewService(storage)
	service.Run()
}
