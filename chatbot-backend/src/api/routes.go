package api

import (
	"time"

	"github.com/gin-contrib/cors"
)

func (s *Server)RegisterRoutes(){
	s.setupMiddleware()

	versionOne  := s.router.Group("/api/v1")
	versionOne.POST("/push-message", s.PushMessage)
	versionOne.POST("/chat", s.ChatWithAgent)

}


func (s *Server) setupMiddleware() {
	s.router.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{
				"http://localhost:5173",        // local dev
        		"http://160.191.162.38:5173",   // backend IP for testing on other devices
			},
			AllowMethods:     []string{"POST", "GET", "OPTIONS"},
			AllowHeaders:     []string{"Content-Type", "Authorization", "X-API-Key"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		},
	))
	s.router.Use(s.APIKeyAuth(),)
}