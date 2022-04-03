package clother

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
}

func (s *Server) Start(addr string) error {
	return s.Engine.Run(addr)
}
