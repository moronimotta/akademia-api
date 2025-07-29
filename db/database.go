package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// GORM:
type GormDatabase struct {
	DB *gorm.DB
}

func (g *GormDatabase) GetDB() *gorm.DB {
	return g.DB
}

// MONGODB:
type MongoDatabase struct {
	DB *mongo.Database
}

func (m *MongoDatabase) GetDB() *mongo.Database {
	return m.DB
}
