package main

import (
	"http-demo/internal/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/create", handler.CreateUser)
	http.HandleFunc("/get", handler.GetUser)
	http.HandleFunc("/delete", handler.DeleteUser)
	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
