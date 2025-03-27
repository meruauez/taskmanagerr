package main

import (
	"log"
	"net/http"

	"taskmanager/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Task Routes
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", handlers.GetTaskByID).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	// User Routes
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")

	// Project Routes
	r.HandleFunc("/projects", handlers.GetProjects).Methods("GET")
	r.HandleFunc("/projects", handlers.CreateProject).Methods("POST")

	log.Println("Task Manager API is running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
