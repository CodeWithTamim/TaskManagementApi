package dto

import "github.com/CodeWithTamim/TaskManagementApi/models"

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateTaskResponse struct {
	Task []models.Task `json:"tasks"`
}

type GetTaskResponse struct {
	Task []models.Task `json:"tasks"`
}

type GetTasksResponse struct {
	Task []models.Task `json:"tasks"`
}
