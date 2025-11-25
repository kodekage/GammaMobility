package utils

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/kodekage/gamma_mobility/internal/logger"
	"github.com/redis/go-redis/v9"
)

func WriteResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func RedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return redisClient
}

func SqlClient() *pgxpool.Pool {
	EnvironmentSetup()
	var ctx = context.Background()

	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Fatal("Unable to Connect to DB")
	}

	// Verify the connection
	if err := pool.Ping(ctx); err != nil {
		log.Fatal("Unable to ping database:", err)
	}

	return pool
}

func EnvironmentSetup() {
	if err := godotenv.Load(); err != nil {
		logger.Warn("No .env file found, proceeding with system environment variables")
	}

	logger.Info("Environment variables loaded")
}
