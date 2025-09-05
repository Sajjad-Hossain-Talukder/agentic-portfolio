package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) APIKeyAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        key := c.GetHeader("X-API-Key")
        if key == "" || key != s.Config.XApiKey {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "Unauthorized: invalid API key",
            })
            return
        }
        c.Next()
    }
}