package service

import (
	"floweys_app/features/tasks"
)

type TaskService struct {
	taskData tasks.TaskDataInterface
}

func (t TaskService) Add(input tasks.TaskCore) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskService) Edit(id uint, input tasks.TaskCore) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskService) Remove(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskService) ReadAll() ([]tasks.TaskCore, error) {
	//TODO implement me
	panic("implement me")
}

func New(repo tasks.TaskDataInterface) tasks.TaskServiceInterface {
	return &TaskService{
		taskData: repo,
	}
}
