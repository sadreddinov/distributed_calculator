package psql

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/models"
)

type TaskPostgres struct {
	db *pgxpool.Pool
}

func NewTaskPostgres(db *pgxpool.Pool) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (r *TaskPostgres) GetTask(id string) (models.Expression, error) {
	var expression models.Expression
	query := fmt.Sprintf("SELECT id, expression, result, created_at, solved_at, work_state, computing_resource_id, user_id FROM %s WHERE work_state = $1 LIMIT 1", expressionTable)
	row := r.db.QueryRow(context.TODO(), query, "in_queue")
	if err := row.Scan(&expression.Id, &expression.Expr, &expression.Result, &expression.Created_at, &expression.Solved_at, &expression.Work_state, &expression.ComputingResourceId, &expression.User_id); err !=nil {
		return models.Expression{}, err
	}
	uuid, _ :=uuid.Parse(id)
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	
	setValues = append(setValues, fmt.Sprintf("computing_resource_id=$%d", argId))
	args = append(args, uuid)
	argId++

	setValues = append(setValues, fmt.Sprintf("work_state=$%d", argId))
	args = append(args, "calculating")
	argId++
	setQuery := strings.Join(setValues, ",")
	args = append(args, expression.Id)
	
	updateQuery := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", expressionTable, setQuery, argId)
	_, err := r.db.Exec(context.TODO(), updateQuery, args...)
	if err != nil {
		return models.Expression{}, err
	}
	return expression, nil
}

func (r *TaskPostgres) PostResult(expression models.Expression) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	
	setValues = append(setValues, fmt.Sprintf("result=$%d", argId))
	args = append(args, expression.Result)
	argId++

	setValues = append(setValues, fmt.Sprintf("work_state=$%d", argId))
	args = append(args, "solved")
	argId++
	
	solved_time, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	setValues = append(setValues, fmt.Sprintf("solved_at=$%d", argId))
	args = append(args, solved_time)
	argId++
	setQuery := strings.Join(setValues, ",")
	args = append(args, expression.Id)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", expressionTable, setQuery, argId)
	_, err := r.db.Exec(context.TODO(), query, args...)
	if err != nil {
		return err
	}
	return nil
}