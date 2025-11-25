package utils

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
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

// func SqlClient() *sqlx.DB {
// 	sqlClient, err := sqlx.Open("mysql", "root:rootpw@/banking")
// 	if err != nil {
// 		logger.Error(err.Error())
// 		panic(err)
// 	}
// 	// See "Important settings" section.
// 	sqlClient.SetConnMaxLifetime(time.Minute * 3)
// 	sqlClient.SetMaxOpenConns(10)
// 	sqlClient.SetMaxIdleConns(10)

// 	return sqlClient
// }
