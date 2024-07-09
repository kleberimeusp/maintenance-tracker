package controllers

import (
	"encoding/json"
	"maintenance-tracker/models"
	"maintenance-tracker/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TaskController struct {
	taskService *services.TaskService
}

func NewTaskController(taskService *services.TaskService) *TaskController {
	return &TaskController{taskService}
}

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.taskService.CreateTask(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *TaskController) GetTasksByUser(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	tasks, err := c.taskService.GetTasksByUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func (c *TaskController) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := c.taskService.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func (c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskID, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := c.taskService.DeleteTask(taskID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
