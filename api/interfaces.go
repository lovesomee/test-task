package api

import "skillsrock-test/models"

type TasksService interface {
	AddTask(task models.Tasks) error
	UpdateTask(task models.Tasks) error
	DeleteTask(id int) error
}
