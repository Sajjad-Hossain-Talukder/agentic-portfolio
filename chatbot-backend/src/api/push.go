package api

import (
	"errors"
	"net/http"

	"github.com/Sajjad-Hossain-Talukder/agentic-portfolio/chatbot-backend/src/services"
	"github.com/gin-gonic/gin"
)

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(), 
	}
}

type PushRequest struct {
	Message string `json:"message" binding:"required"`
}


func (s *Server) PushMessage(ctx *gin.Context){
	var req PushRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("invalid request")))
		return
	}

	
	err := services.PushService(req.Message, s.Config.PushoverUser, s.Config.PushoverToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"recorded": "ok"}) 
}


