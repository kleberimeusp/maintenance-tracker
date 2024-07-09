package services

import (
	"fmt"
	"maintenance-tracker/models"
	"maintenance-tracker/repositories"
	"maintenance-tracker/utils"
)

type TaskService struct {
	taskRepo *repositories.TaskRepository
	userRepo *repositories.UserRepository
	mailer   utils.Mailer
}

func NewTaskService(taskRepo *repositories.TaskRepository, userRepo *repositories.UserRepository, mailer utils.Mailer) *TaskService {
	return &TaskService{taskRepo, userRepo, mailer}
}

func (s *TaskService) CreateTask(task models.Task) error {
	err := s.taskRepo.CreateTask(task)
	if err != nil {
		return err
	}
	user, err := s.userRepo.GetUserById(task.UserID)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("O técnico %s realizou a tarefa %s na data %s", user.Username, task.Summary, task.Date)
	go s.mailer.SendMail("manager@example.com", "Nova Tarefa Realizada", message)
	return nil
}

func (s *TaskService) GetTasksByUser(userID int) ([]models.Task, error) {
	return s.taskRepo.GetTasksByUser(userID)
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.taskRepo.GetAllTasks()
}

func (s *TaskService) DeleteTask(taskID int) error {
	return s.taskRepo.DeleteTask(taskID)
}
