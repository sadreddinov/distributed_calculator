package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/pgtype"
)

type Expression struct {
	Id         uuid.UUID `json:"id" db:"id"`
	Expr       string    `json:"expression" db:"expression"`
	Result     string    `json:"result" db:"result"`
	Work_state string    `json:"work_state" db:"work_state"`
	Created_at pgtype.Timestamptz `json:"created_at" db:"created_at"`
	Solved_at  pgtype.Timestamptz `json:"solved_at" db:"solved_at"`
	ComputingResourceId uuid.UUID `json:"computing_resource_id" db:"computing_resource_id"`
}

type ExpressionToRead struct {
	Id         uuid.UUID `json:"id" db:"id"`
	Expr       string    `json:"expression" db:"expression"`
	Result     string    `json:"result" db:"result"`
	Work_state string    `json:"work_state" db:"work_state"`
	Created_at string `json:"created_at" db:"created_at"`
	Solved_at  string `json:"solved_at" db:"solved_at" `
}
