package postgresRepository

import (
	"akademia-api/db"
	"akademia-api/repositories"
)

type PostgresRepository struct {
	PgPosts
	PgClasses
	PgCourses
}

func NewPostgresRepository(db db.Database) repositories.AkademiaRepository {
	return PostgresRepository{
		PgPosts:   NewPgPostsRepository(db),
		PgClasses: NewPgClassesRepository(db),
		PgCourses: NewPgCoursesRepository(db),
	}
}
