package handler

import "floweys_app/features/tasks"

type TaskHandler struct {
	taskService tasks.TaskServiceInterface
}

func New(service tasks.TaskServiceInterface) *TaskHandler {
	return &TaskHandler{
		taskService: service,
	}
}
