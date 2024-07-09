package main

import (
	"log"
	"maintenance-tracker/config"
	"maintenance-tracker/controllers"
	_ "maintenance-tracker/docs" // Import the generated docs
	"maintenance-tracker/repositories"
	"maintenance-tracker/services"
	"maintenance-tracker/utils"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Maintenance Tracker API
// @version 1.0
// @description This is a sample server for managing maintenance tasks.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

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

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
