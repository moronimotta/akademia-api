package db

import (
	"akademia-api/entities"
	"context"
	"log"
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
		log.Fatalf("❌ Error connecting to PostgreSQL: %v", err)
	}

	// Rodar migrations se necessário
	if err := db.AutoMigrate(
		&entities.Posts{},
		// &entities.Content{},
		&entities.Classes{},
		&entities.Courses{},
	); err != nil {
		log.Fatalf("❌ Error migrating PostgreSQL: %v", err)
	}

	// MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURL := os.Getenv("MONGODB_URL")
	mongoDBName := os.Getenv("MONGODB_NAME")

	clientOpts := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("❌ Error creating MongoDB client: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("❌ Error pinging MongoDB: %v", err)
	}

	mongoDB := client.Database(mongoDBName)
	if mongoDB == nil {
		log.Fatalf("❌ Could not select MongoDB database '%s'", mongoDBName)
	}

	// check if the collection exists, if not, create it
	collectionName := os.Getenv("MONGODB_COLLECTION_NAME")
	if collectionName == "" {
		log.Fatalf("❌ MONGODB_COLLECTION_NAME environment variable is not set")
		if err := mongoDB.CreateCollection(ctx, collectionName); err != nil {
			if !strings.Contains(err.Error(), "already exists") {
				log.Fatalf("❌ Error creating MongoDB collection '%s': %v", collectionName, err)
			}
		}
	}
	log.Printf("✅ MongoDB collection '%s' is ready", collectionName)

	log.Println("✅ MongoDB connected!")
	log.Println("✅ PostgreSQL connected!")

	return &databaseImpl{
		sql:   db,
		mongo: mongoDB,
	}, nil
}
