package service

import (
	"lobo.tech/task/model"
	"lobo.tech/task/repository"
)

func NewTaskService() *TaskService {

	db := repository.NewDB()
	taskRespository := &repository.TaskRepository{DB: db}
	taskService := &TaskService{taskRespository}
	return taskService
}

type TaskService struct {
	TaskRepository *repository.TaskRepository
}

func (service *TaskService) GetByID(id string) (model.Task, error) {

	task, err := service.TaskRepository.GetByID(id)
	return task, err

}
