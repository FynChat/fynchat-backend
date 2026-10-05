package models

import (
	"time"

	"github.com/google/uuid"
)

type Messages struct {
	ID        uuid.UUID `json:"id"`
	ChannelID uuid.UUID `json:"channel_id"`
	SenderID  int64     `json:"sender_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
