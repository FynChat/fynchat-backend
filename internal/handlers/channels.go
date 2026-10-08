package handlers

import (
	"fynchat.thegrigoriyagen.com/internal/models"
	"fynchat.thegrigoriyagen.com/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChannelRequest struct {
	ChannelType string `json:"channel_type" binding:"required"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar_url"`
}

type ChannelResponse struct {
	Name   string `json:"channel_name"`
	Avatar string `json:"channel_avatar"`
}

type ChannelHandler struct {
	channelService *services.ChannelService
}

func ToChannelResponse(channel *models.Channels) ChannelResponse {
	return ChannelResponse{
		Name:   channel.ChannelName,
		Avatar: channel.Avatar,
	}
}

func NewChannelHandler(channelService *services.ChannelService) *ChannelHandler {
	return &ChannelHandler{channelService: channelService}
}

func (h *ChannelHandler) Create(c *gin.Context) {
	ownerID := c.GetInt("userID")

	var input ChannelRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	channel, err := h.channelService.Create(c.Request.Context(), ownerID, input.ChannelType, input.Name, input.Avatar)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ToChannelResponse(channel))
}
