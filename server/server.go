package server

import (
	"akademia-api/handlers"
	logs "akademia-api/utils/logs"

	"akademia-api/db"

	"github.com/gin-gonic/gin"
)

type Server struct {
	app       *gin.Engine
	pgHandler *handlers.DbHttpHandler
}

func NewServer(db db.Database) *Server {
	logs.InitLogging()

	pgHandler, err := handlers.NewDbHttpHandler("postgres", db)
	if err != nil {
		return nil
	}

	return &Server{
		app:       gin.Default(),
		pgHandler: pgHandler,
	}
}
func (s *Server) Start() {

	s.initializeMiddlewares()

	s.initializeAkademiaHttpHandler()

	if err := s.app.Run(":3536"); err != nil {
		panic(err)
	}
}

func (s *Server) initializeAkademiaHttpHandler() {

	s.initPostRoutes()

}
