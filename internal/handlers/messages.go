package handlers

import (
	"fynchat/internal/models"
	"fynchat/internal/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MessageRequest struct {
	Content string `json:"content"`
}

type MessageResponse struct {
	ID        uuid.UUID `json:"id"`
	ChannelID uuid.UUID `json:"channel_id"`
	SenderID  int64     `json:"sender_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type MessageHandler struct {
	messageService *services.MessageService
}

func toMessageResponse(m *models.Messages) MessageResponse {
	return MessageResponse{
		ID:        m.ID,
		ChannelID: m.ChannelID,
		SenderID:  m.SenderID,
		Content:   m.Content,
		CreatedAt: m.CreatedAt,
	}
}

func (h *MessageHandler) Create(c *gin.Context) {
	senderID := int64(c.GetInt("userID"))
	channelIDStr := c.Param("channel_id")

	channelID, err := uuid.Parse(channelIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid channel id"})
		return
	}

	var input MessageRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := h.messageService.Create(c.Request.Context(), senderID, channelID, input.Content)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, toMessageResponse(message))
}
