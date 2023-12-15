// config.go

package config

import (
	"context"
	"errors"
	"fmt"
	"time"

	"hexagon-architecture/internal/infrastructure"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

type config struct {
	App  appConfig  `envconfig:"APP"`
	Otel otelConfig `envconfig:"OTEL"`
	DB   dbConfig   `envconfig:"DB"`
}

type appConfig struct {
	Env             string        `envconfig:"ENV" default:"development" validate:"required,oneof=local development staging production" mod:"no_space,lcase"` // local
	Port            string        `envconfig:"PORT" default:"8006" validate:"required" mod:"no_space"`
	ReadTimeout     time.Duration `envconfig:"READ_TIMEOUT" default:"5s" validate:"required,gt=0"`
	WriteTimeout    time.Duration `envconfig:"WRITE_TIMEOUT" default:"5s" validate:"required,gt=0"`
	GracefulTimeout time.Duration `envconfig:"GRACEFUL_TIMEOUT" default:"10s" validate:"required,gt=0"`
	Host            string        `envconfig:"HOST" validate:"required,url"`
	Name            string        `envconfig:"NAME" validate:"required"`
	Version         string        `envconfig:"VERSION" validate:"required"`
}

type otelConfig struct {
	Host string `envconfig:"GRPC_HOST" validate:"required"`
}

type dbConfig struct {
	Address      string        `envconfig:"ADDRESS" default:"localhost:27018" validate:"required"`
	Name            string        `envconfig:"NAME" default:"store" validate:"required"`
	// User            string        `envconfig:"USER" default:"root" validate:"required"`
	// Password        string        `envconfig:"PASSWORD"`
	SslMode         string        `envconfig:"SSL_MODE" default:"disable"`
	MaxConnOpen     int           `envconfig:"MAX_CONN_OPEN" default:"10" validate:"required,gt=0"`
	MaxConnIdle     int           `envconfig:"MAX_CONN_IDLE" default:"10" validate:"required,gt=0"`
	MaxConnLifetime time.Duration `envconfig:"MAX_CONN_LIFETIME" default:"60s" validate:"required,gt=0"`
}

const envPrefix = "STORE"

func GetConfig() *config {
	var cfg config

	// Load .env file.
	_ = godotenv.Load()

	// Convert env to struct.
	if err := envconfig.Process(envPrefix, &cfg); err != nil {
		return nil
	}

	// Init global log.
	infrastructure.InitSlog()

	return &cfg
}

func NewDB(cfg dbConfig, env string) (*mongo.Database, error) {

	// dbOptions to setup mongodb uri
	var dns string
	switch env {
	case "local", "development":
		dns = fmt.Sprintf("mongodb://%s", cfg.Address)
	default:
	}

	dbOptions := options.Client()
	dbOptions.Monitor = otelmongo.NewMonitor()
	dbOptions.ApplyURI(dns)

	// Prepare dns and open connection.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connection to mongo database
	dbClient, err := mongo.Connect(ctx, dbOptions)
	if err != nil {
		fmt.Println("=========== Failed to connect to database ", err)
		return nil, errors.Join(err)
	}

	fmt.Println("=========== success to connect to database ")

	//tess
	err = dbClient.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("=========== Failed to ping database: ", err)
		return nil, err
	}

	fmt.Println("=========== success ping to connect to database ")

	// Test ping to the database
	// err = dbClient.Ping(ctx, nil)
	// if err != nil {
	// 	fmt.Println("=========== Failed to ping database: ", err)

	// 	// Check if it's a server selection error
	// 	if cmdErr, ok := err.(mongo.CommandError); ok && cmdErr.HasErrorLabel("ServerSelectionTimeout") {
	// 		fmt.Printf("============== Server selection timeout error: %v\n", cmdErr)
	// 		return nil, fmt.Errorf("Server selection timeout: %w", err)
	// 	}

	// 	return nil, err
	// }

	// fmt.Println("=========== success ping to connect to database ")

	return dbClient.Database(cfg.Name), nil

}
