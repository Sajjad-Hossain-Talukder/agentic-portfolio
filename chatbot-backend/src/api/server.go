package api

import (
	chatagent "github.com/Sajjad-Hossain-Talukder/agentic-portfolio/chatbot-backend/src/agent"
	"github.com/Sajjad-Hossain-Talukder/agentic-portfolio/chatbot-backend/src/config"
	"github.com/gin-gonic/gin"
)


type Server struct {
	router 		*gin.Engine
	Config 		*config.Config
	ChatAgent 	*chatagent.ChatAgent
}

func NewServer(config *config.Config, chatAgent *chatagent.ChatAgent) *Server {
	server := Server { 
		Config: config, 
		router: gin.Default(), 
		ChatAgent: chatAgent,
	}
	server.RegisterRoutes()
	return &server
}

func (s *Server) RunWith (address string) error {
	return s.router.Run(address)
}