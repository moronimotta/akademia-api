package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Database interface {
	GetSQLDB() *gorm.DB
	GetMongoDB() *mongo.Database
}

type databaseImpl struct {
	sql   *gorm.DB
	mongo *mongo.Database
}

func (d *databaseImpl) GetSQLDB() *gorm.DB {
	return d.sql
}

func (d *databaseImpl) GetMongoDB() *mongo.Database {
	return d.mongo
}
