package service

import (
	"github.com/google/uuid"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/models"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/storage/psql"
)

type ExpressionService struct {
	repo psql.Expression
}

func NewExpressionService(repo psql.Expression) *ExpressionService {
	return &ExpressionService{repo: repo}
}

func (s *ExpressionService) GetExpression(uuid uuid.UUID) (models.ExpressionToRead, error) {
	return s.repo.GetExpression(uuid)
}

func (s *ExpressionService) GetExpressions(startIndex int, recordPerPage int) ([]models.ExpressionToRead, error) {
	return s.repo.GetExpressions(startIndex, recordPerPage)
}

func (s *ExpressionService) CreateExpression(expression models.Expression) (uuid.UUID, error) {
	return s.repo.CreateExpression(expression)
}
