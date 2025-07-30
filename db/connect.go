package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (Database, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err := db.AutoMigrate(
	// &entities.Posts{},
	// &entities.Content{},
	// &entities.Classes{},
	// &entities.Courses{},
	); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	// Modern MongoDB connection
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err)
	}

	// Test the connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	mongoDB := client.Database(os.Getenv("MONGODB_NAME"))
	if mongoDB == nil {
		log.Fatalf("Error getting MongoDB database")
	}

	// Index creation moved to manual setup or migration script
	// collection := mongoDB.Collection(os.Getenv("MONGODB_COLLECTION_NAME"))
	// _, err = collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
	// 	Keys:    bson.D{{Key: "user_id", Value: 1}},
	// 	Options: options.Index().SetUnique(true),
	// })
	// if err != nil {
	// 	log.Printf("Warning: Error creating index on MongoDB collection: %v", err)
	// }

	return &databaseImpl{
		sql:   db,
		mongo: mongoDB,
	}, nil
}
