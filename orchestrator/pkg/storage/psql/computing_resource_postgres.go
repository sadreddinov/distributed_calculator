package psql

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/models"
)

type ComputingResourcePostgres struct {
	db *pgxpool.Pool
}

func NewComputingResourcePostgres(db *pgxpool.Pool) *ComputingResourcePostgres {
	return &ComputingResourcePostgres{db: db}
}

func (r *ComputingResourcePostgres) CreateComputingResource(agent models.ComputingResource) (uuid.UUID, error) {
	var old_agent models.ComputingResource
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", computingResourceTable)
	row := r.db.QueryRow(context.TODO(), query, agent.Id)
	if err := row.Scan(&old_agent.Id, &old_agent.Work_state, &old_agent.LastPing_at); err == pgx.ErrNoRows {
		var id uuid.UUID
		agent.LastPing_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		query := fmt.Sprintf("INSERT INTO %s (id, work_state,last_ping_at) VALUES ($1, $2, $3) RETURNING id", computingResourceTable)
		row := r.db.QueryRow(context.TODO(), query, agent.Id, agent.Work_state, agent.LastPing_at)
		if err := row.Scan(&id); err != nil {
			return uuid.UUID{}, err
		}
		return id, nil
	} else if err != nil {
		return uuid.UUID{}, nil
	} else {
		setValues := make([]string, 0)
		args := make([]interface{}, 0)
		argId := 1

		setValues = append(setValues, fmt.Sprintf("work_state=$%d", argId))
		args = append(args, "is_working")
		argId++

		ping_time, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		setValues = append(setValues, fmt.Sprintf("last_ping_at=$%d", argId))
		args = append(args, ping_time)
		argId++

		setQuery := strings.Join(setValues, ",")
		args = append(args, agent.Id)
		query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", computingResourceTable, setQuery, argId)
		_, err := r.db.Exec(context.TODO(), query, args...)
		if err != nil {
			return uuid.UUID{}, err
		}
		return agent.Id, nil
	}
}

func (r *ComputingResourcePostgres) GetComputingResources() ([]models.ComputingResource, error) {
	var agents []models.ComputingResource
	query := fmt.Sprintf("SELECT * FROM %s", computingResourceTable)
	rows, err := r.db.Query(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var agent models.ComputingResource
		err := rows.Scan(&agent.Id, &agent.Work_state, &agent.LastPing_at)
		if err != nil {
			return nil, err
		}
		if agent.LastPing_at.Add(10*time.Minute).After(time.Now()) {
			agent.Work_state = "lost_connection"
			setQuery := fmt.Sprintf("work_state=$%d", 1)
			query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $2", computingResourceTable, setQuery)
			_, err = r.db.Exec(context.TODO(), query, agent.Work_state, agent.Id)
		}
		agents = append(agents, agent)
	}
	return agents, nil
}

func (r *ComputingResourcePostgres) UpdateComputingResource(agent models.ComputingResource) error {
	ping_time, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	setQuery := fmt.Sprintf("last_ping_at=$%d", 1)

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $2", computingResourceTable, setQuery)
	_, err := r.db.Exec(context.TODO(), query, ping_time, agent.Id)
	return err
}

func (r *ComputingResourcePostgres) ShutdownComputingResource(agent models.ComputingResource) (error) {
	setQuery := fmt.Sprintf("work_state=$%d", 1)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $2", computingResourceTable, setQuery)
	_, err := r.db.Exec(context.TODO(), query, agent.Work_state, agent.Id)
	query = fmt.Sprintf("SELECT id FROM %s WHERE computing_resource_id = $1 ", expressionTable)
	rows, err := r.db.Query(context.TODO(), query, agent.Id)
	if err != nil {
		return err
	}
	for rows.Next() {
		var exp models.Expression
		err := rows.Scan(&exp.Id)
		if err != nil {
			return err
		}
		setValues := make([]string, 0)
		args := make([]interface{}, 0)
		argId := 1

		setValues = append(setValues, fmt.Sprintf("work_state=$%d", argId))
		args = append(args, "in_queue")
		argId++

		setValues = append(setValues, fmt.Sprintf("computing_resource_id=$%d", argId))
		id, _ :=uuid.Parse("00000000-0000-0000-0000-000000000000")
		args = append(args, id)
		argId++

		setQuery := strings.Join(setValues, ",")
		args = append(args, exp.Id)
		query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", expressionTable, setQuery, argId)
		_, err = r.db.Exec(context.TODO(), query, args...)
		if err != nil {
			return err
		}
	}
	return err
}