package main

import (
	"context"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/service"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/storage/psql"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/transport"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/transport/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title           Distributed Calculator
// @version         1.0
// @description     Это простой распределенный калькулятор.

// @contact.name   Фарид
// @contact.url    https://t.me/M00nfI0wer
// @contact.email  faridsadreddinov@yandex.ru

// @host      localhost:8080
// @BasePath  /

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initconfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(filepath.Join("..\\", ".env")); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := psql.NewPostgresDb(psql.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DbName:   viper.GetString("db.dbname"),
		Password: os.Getenv("DB_PASSWORD"),
		Sslmode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := psql.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(transport.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error ocured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down", err.Error())
	}
	db.Close()
}

func initconfig() error {
	viper.AddConfigPath("..\\configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
