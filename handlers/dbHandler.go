package handlers

import (
	"akademia-api/db"
	usescases "akademia-api/usecases/db"
)

type DbHttpHandler struct {
	usescases.DbUsecase
}

func NewDbHttpHandler(db db.Database) (*DbHttpHandler, error) {
	return &DbHttpHandler{
		*usescases.NewDbUsecase(db),
	}, nil
}
