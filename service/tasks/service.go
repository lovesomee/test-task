package tasks

import (
	"skillsrock-test/models"
)

type IRepository interface {
	InsertTasks(tasks models.Tasks) error
	UpdateTasks(tasks models.Tasks) error
	DeleteTasks(id int) error
	SelectTasks() ([]models.Tasks, error)
}

type Service struct {
	tasksRepository IRepository
}

func NewService(tasksRepository IRepository) *Service {
	return &Service{tasksRepository: tasksRepository}
}

func (s *Service) AddTask(task models.Tasks) error {
	if err := s.tasksRepository.InsertTasks(task); err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateTask(task models.Tasks) error {
	if err := s.tasksRepository.UpdateTasks(task); err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteTask(id int) error {
	if err := s.tasksRepository.DeleteTasks(id); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetTask() ([]models.Tasks, error) {
	return s.tasksRepository.SelectTasks()
}
