package service

import (
	"time"

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

func (service *TaskService) Save(task model.Task) (model.Task, error) {

	taskID, err := service.TaskRepository.Save(task)
	if err != nil {
		return model.Task{}, err
	}
	tasknew, err := service.TaskRepository.GetByID(taskID)
	return tasknew, err

}

func (service *TaskService) Update(task model.Task) (model.Task, error) {

	qtdUpdate, err := service.TaskRepository.Update(task)
	if err != nil {
		return model.Task{}, err
	}
	if qtdUpdate > 1 {

		//fururamente trasaction - begin commit
		panic("Quantidate de update superior " + string(qtdUpdate))
	}

	tasknew, err := service.TaskRepository.GetByID(task.ID)

	return tasknew, err
}

func (service *TaskService) Delete(taskID string) (bool, error) {

	qtdDeleted, err := service.TaskRepository.Delete(taskID)
	if err != nil {
		return false, err
	}
	if qtdDeleted > 1 {
		//fururamente trasaction - begin commit
		panic("Quantidate de update superior " + string(qtdDeleted))
	}

	return true, err
}

func (service *TaskService) GetByPeriod(start, finish string) ([]model.Task, error) {

	layoutISO := "2006-01-02"
	date_start, err_start := time.Parse(layoutISO, start)
	date_finish, err_finish := time.Parse(layoutISO, finish)

	if err_start != nil || err_finish != nil {
		return nil, err_start
	}

	var tasks []model.Task = []model.Task{}

	tasks, err := service.TaskRepository.GetByPeriod(date_start, date_finish)

	return tasks, err

}
