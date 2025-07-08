package handlers

import (
	"akademia-api/db"
	usescases "akademia-api/usecases/db"
	"errors"
)

type DbHttpHandler struct {
	usescases.DbUsecase
}

func NewDbHttpHandler(dbName string, db db.Database) (*DbHttpHandler, error) {
	var usecaseInput usescases.DbUsecase

	switch dbName {
	case "postgres":
		usecaseInput = *usescases.NewPgUsecase(db)
	default:
		return nil, errors.New("unsupported database")
	}
	return &DbHttpHandler{
		usecaseInput,
	}, nil
}
