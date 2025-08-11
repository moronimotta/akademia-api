package server

import (
	"context"
	"log/slog"
	"os"

	"akademia-api/db"
	logs "akademia-api/utils/logs"

	"github.com/redis/go-redis/v9"
)

type RedisServer struct {
	db            db.Database
	connectionUrl string
	queueName     string
	topicName     string
	Client        *redis.Client // Export the client
}

func NewRedisServer(db db.Database) *RedisServer {
	logs.InitLogging()
	// TODO: Change Redis PORT
	connectionUrl := os.Getenv("REDIS_URL")
	if connectionUrl == "" {
		connectionUrl = "redis://localhost:6379" // Default fallback
	}

	return &RedisServer{
		db:            db,
		connectionUrl: connectionUrl,
	}
}

func (s *RedisServer) Start() *redis.Client {
	opt, err := redis.ParseURL(s.connectionUrl)
	if err != nil {
		slog.Error("Failed to parse Redis URL")
		panic(err)
	}

	rdb := redis.NewClient(opt)

	// Test the connection
	ctx := context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		slog.Error("Failed to connect to Redis: %v", err)
		panic("Failed to connect to Redis: " + err.Error())
	}

	s.Client = rdb
	return rdb
}

func (s *RedisServer) Close() error {
	if s.Client != nil {
		slog.Info("Closing Redis connection")
		return s.Client.Close()
	}
	return nil
}
