package main

import (
	"log"
	"maintenance-tracker/config"
	"maintenance-tracker/controllers"
	"maintenance-tracker/repositories"
	"maintenance-tracker/services"
	"maintenance-tracker/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := config.GetDB()
	mailer := utils.NewMailer()

	userRepo := repositories.NewUserRepository(db)
	taskRepo := repositories.NewTaskRepository(db)

	userService := services.NewUserService(userRepo)
	taskService := services.NewTaskService(taskRepo, userRepo, *mailer)

	userController := controllers.NewUserController(userService)
	taskController := controllers.NewTaskController(taskService)

	r := mux.NewRouter()

	r.HandleFunc("/users", userController.CreateUser).Methods("POST")
	r.HandleFunc("/tasks", taskController.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", taskController.GetTasksByUser).Methods("GET").Queries("user_id", "{user_id}")
	r.HandleFunc("/tasks/all", taskController.GetAllTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", taskController.DeleteTask).Methods("DELETE")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
