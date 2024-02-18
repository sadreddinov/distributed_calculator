package service

import (
	"github.com/google/uuid"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/models"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/storage/psql"
)

type Expression interface {
	GetExpression(uuid uuid.UUID) (models.ExpressionToRead, error)
	GetExpressions(startIndex int, recordPerPage int) ([]models.ExpressionToRead, error)
	CreateExpression(expression models.Expression) (uuid.UUID, error)
}

type ComputingResource interface {
	CreateComputingResource(agent models.ComputingResource) (uuid.UUID, error)
	GetComputingResources() ([]models.ComputingResource, error)
	UpdateComputingResource(agent models.ComputingResource) (error)
	ShutdownComputingResource(agent models.ComputingResource) (error)
}

type Task interface {
	GetTask(id string) (models.Expression, error)
	PostResult(expression models.Expression) error 
}

type Service struct {
	Expression
	ComputingResource
	Task
}

func NewService(repos *psql.Repository) *Service {
	return &Service{
		Expression:        NewExpressionService(repos.Expression),
		ComputingResource: NewComputingResourceService(repos.ComputingResource),
		Task: NewTaskService(repos.Task),
	}
}
