package psql

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Expression interface {
	GetExpression(uuid uuid.UUID, userId int) (models.ExpressionToRead, error)
	GetExpressions(startIndex int, recordPerPage int, userId int) ([]models.ExpressionToRead, error)
	CreateExpression(expression models.ExpressionFromUser, userId int) (uuid.UUID, error)
}

type ComputingResource interface {
	CreateComputingResource(agent models.ComputingResource) (uuid.UUID, error)
	GetComputingResources() ([]models.ComputingResource, error)
	UpdateComputingResource (agent models.ComputingResource) (error)
	ShutdownComputingResource (agent models.ComputingResource) (error)
}

type Task interface {
	GetTask(id string) (models.Expression, error)
	PostResult(expression models.Expression) error
}

type Repository struct {
	Authorization
	Expression
	ComputingResource
	Task
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Expression:        NewExpressionPostgres(db),
		ComputingResource: NewComputingResourcePostgres(db),
		Task: NewTaskPostgres(db),
	}
}
