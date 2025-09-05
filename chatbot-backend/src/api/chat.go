package api

import (
	"log"
	"net/http"

	chatagent "github.com/Sajjad-Hossain-Talukder/agentic-portfolio/chatbot-backend/src/agent"
	"github.com/gin-gonic/gin"
)


type ChatRequest struct {
	Message string        `json:"message"`
	History []chatagent.ChatMessage `json:"history"`
}

type ChatResponse struct {
	Reply string `json:"reply"`
}

func (s *Server) ChatWithAgent(ctx *gin.Context) {
    var req ChatRequest
    if err := ctx.BindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    reply, err := s.ChatAgent.Chat(req.Message, req.History)
    if err != nil {
        log.Println("Chat error:", err)
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "chat failed"})
        return
    }
    ctx.JSON(http.StatusOK, ChatResponse{Reply: reply})
    
}
