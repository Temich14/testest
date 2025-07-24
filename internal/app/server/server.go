package server

import (
	"awesomeProject2/internal/adapters/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	api *gin.Engine
}

func NewServer(getHandler handlers.GetHandler) *Server {
	api := gin.Default()

	profileGroup := api.Group("/profiles")

	profileGroup.GET("/:id", getHandler.GetByID)

	return &Server{api: api}
}

func (s *Server) Run() error {
	err := s.api.Run(":8081")
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
