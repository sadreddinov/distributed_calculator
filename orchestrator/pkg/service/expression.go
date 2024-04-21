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

func (s *ExpressionService) GetExpression(uuid uuid.UUID, userId int) (models.ExpressionToRead, error) {
	return s.repo.GetExpression(uuid, userId)
}

func (s *ExpressionService) GetExpressions(startIndex int, recordPerPage int, userId int) ([]models.ExpressionToRead, error) {
	return s.repo.GetExpressions(startIndex, recordPerPage, userId)
}

func (s *ExpressionService) CreateExpression(expression models.ExpressionFromUser, userId int) (uuid.UUID, error) {
	return s.repo.CreateExpression(expression, userId)
}
