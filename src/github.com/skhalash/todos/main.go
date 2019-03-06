package main

import "github.com/skhalash/todos/service"

func main() {
	service := service.NewService()
	service.Run()
}
