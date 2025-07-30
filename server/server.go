package server

import (
	"akademia-api/handlers"
	logs "akademia-api/utils/logs"

	"akademia-api/db"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	app         *gin.Engine
	dbHandler   *handlers.DbHttpHandler
	redisClient *redis.Client
}

func NewServer(db db.Database, redisClient *redis.Client) *Server {
	logs.InitLogging()

	dbHandler, err := handlers.NewDbHttpHandler(db)
	if err != nil {
		return nil
	}

	return &Server{
		app:         gin.Default(),
		dbHandler:   dbHandler,
		redisClient: redisClient,
	}
}
func (s *Server) Start() {

	s.initializeMiddlewares()

	s.initializeAkademiaHttpHandler()

	// s.app.POST("/content/upload-url", {

	// })

	if err := s.app.Run(":3536"); err != nil {
		panic(err)
	}
}

func (s *Server) initializeAkademiaHttpHandler() {

	s.initPostRoutes()
	s.initUserProgressRoutes()
	s.initCoursesRoutes()
	s.initClassesRoutes()
}
