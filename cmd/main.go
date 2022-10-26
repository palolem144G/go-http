package main

import (
	"http-demo/internal/handler"
	"http-demo/internal/repository"
	"http-demo/internal/service"
	"log"
	"net/http"
)

func main() {
	userStore := repository.NewInMemoryUserStore()
	service := service.NewService(userStore)
	handlers := handler.NewUserHandler(service)

	// map handlers
	http.HandleFunc("/create", handlers.Create)
	http.HandleFunc("/get", handlers.GetUser)
	http.HandleFunc("/pwd", handlers.ChangePassword)
	http.HandleFunc("/delete", handlers.DeleteUser)

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
