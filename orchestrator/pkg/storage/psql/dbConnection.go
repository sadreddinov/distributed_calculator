package psql

import (
	"context"
	"fmt"
	
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	expressionTable = "expression"
	computingResourceTable = "computing_resource"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	Sslmode  string
}

func NewPostgresDb(cfg Config) (*pgxpool.Pool, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",cfg.Username,cfg.Password,cfg.Host,cfg.Port,cfg.DbName,)
	pool, err := pgxpool.New(context.TODO(), connectionString)
	if err != nil {
		return nil, err
	}
	return pool, nil
}