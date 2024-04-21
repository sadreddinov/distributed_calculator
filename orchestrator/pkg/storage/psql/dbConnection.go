package psql

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	usersTable      = "users"
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

func Migrate(cfg Config) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",cfg.Username,cfg.Password,cfg.Host,cfg.Port,cfg.DbName,cfg.Sslmode)
	os.Getwd()
	m, err := migrate.New(
		"file://" + "./schema",
		connectionString)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != migrate.ErrNoChange && err != nil{
		log.Fatal(err)
	}
}