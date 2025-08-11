package db

import (
	"akademia-api/entities"
	"context"
	"errors"
	"log/slog"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (Database, error) {
	// PostgreSQL
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Rodar migrations se necessário
	if err := db.AutoMigrate(
		&entities.Posts{},
		// &entities.Content{},
		&entities.Classes{},
		&entities.Courses{},
	); err != nil {
		return nil, err
	}

	// MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURL := os.Getenv("MONGODB_URL")
	mongoDBName := os.Getenv("MONGODB_NAME")

	clientOpts := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	mongoDB := client.Database(mongoDBName)
	if mongoDB == nil {
		return nil, errors.New("could not select MongoDB database")
	}

	// check if the collection exists, if not, create it
	collectionName := os.Getenv("MONGODB_COLLECTION_NAME")
	if collectionName == "" {
		return nil, errors.New("MONGODB_COLLECTION_NAME environment variable is not set")
	}
	if err := mongoDB.CreateCollection(ctx, collectionName); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			return nil, err
		}
	}

	slog.Info("MongoDB collection '%s' is ready", collectionName)
	slog.Info("MongoDB connected!")
	slog.Info("PostgreSQL connected!")

	return &databaseImpl{
		sql:   db,
		mongo: mongoDB,
	}, nil
}
