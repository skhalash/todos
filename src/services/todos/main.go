package main

import "todos/service"

func main() {
	service := service.NewService()
	service.Run()
}
