package main

import "service"

func main() {
	service := service.NewService()
	service.Run()
}
