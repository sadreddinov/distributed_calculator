package psql

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/models"
)

type ExpressionPostgres struct {
	db *pgxpool.Pool
}

func NewExpressionPostgres(db *pgxpool.Pool) *ExpressionPostgres {
	return &ExpressionPostgres{db: db}
}

func (r *ExpressionPostgres) GetExpression(uuid uuid.UUID, userId int) (models.ExpressionToRead, error) {
	var expression models.Expression

	query := fmt.Sprintf("SELECT id, expression, result, created_at, solved_at, work_state FROM %s WHERE id = $1 AND user_id = $2", expressionTable)
	row := r.db.QueryRow(context.TODO(), query, uuid, userId)
	if err := row.Scan(&expression.Id, &expression.Expr, &expression.Result, &expression.Created_at, &expression.Solved_at, &expression.Work_state); err != nil {
		return models.ExpressionToRead{}, err
	}
	return toRead(expression), nil
}

func (r *ExpressionPostgres) GetExpressions(startIndex int, recordPerPage int, userId int) ([]models.ExpressionToRead, error) {
	var expressions []models.ExpressionToRead
	query := fmt.Sprintf("SELECT id, expression, result, created_at, solved_at, work_state FROM %s WHERE user_id = $1 LIMIT $2 OFFSET $3", expressionTable)
	rows, err := r.db.Query(context.TODO(), query, userId, recordPerPage, startIndex)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var exp models.Expression
		err := rows.Scan(&exp.Id, &exp.Expr, &exp.Result, &exp.Created_at, &exp.Solved_at, &exp.Work_state)
		if err != nil {
			return nil, err
		}
		expressions = append(expressions, toRead(exp))
	}
	return expressions, nil
}
func (r *ExpressionPostgres) CreateExpression(expression models.ExpressionFromUser, userId int) (uuid.UUID, error) {
	id := uuid.New()
	var expr models.Expression
	expr.Id = id
	expr.Expr = expression.Expr
	expr.Created_at.Time, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	expr.Work_state = "in_queue"
	query := fmt.Sprintf("INSERT INTO %s (id, expression, created_at, work_state, user_id) VALUES ($1, $2, $3, $4, $5) RETURNING id", expressionTable)
	row := r.db.QueryRow(context.TODO(), query, expr.Id, expr.Expr, expr.Created_at.Time, expr.Work_state, userId)
	if err := row.Scan(&id); err != nil {
		return uuid.UUID{}, err
	}
	return id, nil
}

func toRead(exp models.Expression) models.ExpressionToRead {
	var expressionToRead models.ExpressionToRead
	expressionToRead.Id = exp.Id
	expressionToRead.Expr = exp.Expr
	expressionToRead.Result = exp.Result
	expressionToRead.Work_state = exp.Work_state
	expressionToRead.Created_at = exp.Created_at.Time.String()
	expressionToRead.Solved_at = exp.Solved_at.Time.String()
	if expressionToRead.Solved_at == "0001-01-01 00:00:00 +0000 UTC" {
		expressionToRead.Solved_at = "soon"
	}
	return expressionToRead
}
