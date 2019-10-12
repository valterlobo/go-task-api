package model

import "time"

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	DateStart   time.Time `json:"date_start"`
	DateStop    time.Time `json:"date_stop"`
	DateCreate  time.Time `json:"date_create"`
}
