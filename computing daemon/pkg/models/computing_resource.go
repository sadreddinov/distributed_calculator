package models

import (
	"time"

	"github.com/google/uuid"
)

type Agent struct {
	Id       uuid.UUID    `json:"id" db:"id"`
	Work_state   string `json:"work_state" db:"work_state"`
	LastPing_at time.Time `json:"last_ping_at" db:"last_ping_at"`
}