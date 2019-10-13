package repository

import (
	"database/sql"
	"time"

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

func (repository *TaskRepository) Save(task model.Task) (string, error) {

	sqlStatement :=
		`
	INSERT INTO public.task(
		id, title, description, done, date_start, date_stop, date_create)
		VALUES (uuid_generate_v4(),  $1 , $2 , false , null , null , now() ) 
	RETURNING id`

	id := ""
	err := repository.DB.QueryRow(sqlStatement, task.Title, task.Description).Scan(&id)

	if err != nil {
		return id, err
	}
	return id, nil

}

func (repository *TaskRepository) Update(task model.Task) (int, error) {

	sqlStatement := `
			UPDATE task
			SET title=$1, description=$2, done=$3, date_start=$4, date_stop=$5 
			WHERE id=$6 `

	res, err := repository.DB.Exec(sqlStatement, task.Title, task.Description, task.Done, task.DateStart, task.DateStop, task.ID)

	rowsUpdated, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}

	return int(rowsUpdated), nil
}

func (repository *TaskRepository) Delete(id string) (int, error) {

	res, err := repository.DB.Exec("DELETE FROM task WHERE id=$1", id)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil

}

func (repository *TaskRepository) GetByPeriod(date_start, date_finish time.Time) ([]model.Task, error) {

	var tasks []model.Task = []model.Task{}

	rows, err := repository.DB.Query(`SELECT 
							id, 
							title, 
	                        description, 
							done,
							date_start,
							date_stop, 
							date_create
							FROM task  
							WHERE 
							DATE(date_create) >= $1 
							and DATE(date_create) <= $2 `, date_start, date_finish)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var task model.Task

	for rows.Next() {
		rows.Scan(&task.ID, &task.Title, &task.Description,
			&task.Done, &task.DateStart, &task.DateStop, &task.DateCreate)
		tasks = append(tasks, task)
	}

	return tasks, nil
}
