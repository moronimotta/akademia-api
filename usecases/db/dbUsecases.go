package usecases

import (
	"akademia-api/db"
	"akademia-api/repositories"
	postgresRepository "akademia-api/repositories/db/postgres"
)

type DbUsecase struct {
	Repository repositories.AkademiaRepository
}

func NewPgUsecase(db db.Database) *DbUsecase {

	repository := postgresRepository.NewPostgresRepository(db)

	return &DbUsecase{
		Repository: repository,
	}
}
