package service

import (
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/models"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/storage/psql"
)

type TaskService struct {
	repo psql.Task
}

func NewTaskService(repo psql.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetTask(id string) (models.Expression, error) {
	return s.repo.GetTask(id)
}

func (s *TaskService) PostResult(expression models.Expression) error {
	return s.repo.PostResult(expression)
}