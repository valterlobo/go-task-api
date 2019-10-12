package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
	"lobo.tech/task/model"
)

type TaskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {

	return &TaskRepository{DB: db}

}

func (repository *TaskRepository) GetByID(id string) (model.Task, error) {

	var task model.Task = model.Task{}

	//id, title, description, done, date_start, date_stop, date_create
	rows, err := repository.DB.Query(`SELECT id, title, 
	                        description, 
							done,
							date_start,
							date_stop , 
							date_create
							FROM task  WHERE id = $1`, id)

	if err != nil {
		return task, err
	}

	if rows.Next() {
		rows.Scan(&task.ID, &task.Title, &task.Description,
			&task.Done, &task.DateStart, &task.DateStop, &task.DateCreate)
	}
	defer rows.Close()

	return task, nil

}

func (repository *TaskRepository) Save(task model.Task) (model.Task, error) {

	//id, title, description, done, date_start, date_stop, date_create
	rows, err := repository.DB.Query(`SELECT id, title, 
	                        description, 
							done,
							date_start,
							date_stop , 
							date_create
							FROM task  WHERE id = $1`, task.ID)

	if err != nil {
		return task, err
	}

	if rows.Next() {
		rows.Scan(&task.ID, &task.Title, &task.Description,
			&task.Done, &task.DateStart, &task.DateStop, &task.DateCreate)
	}
	defer rows.Close()

	return task, nil
}

func (repository *TaskRepository) Update(task model.Task) (model.Task, error) {

	return task, nil
}

func (repository *TaskRepository) Delete(id string) (bool, error) {

	return true, nil

}
