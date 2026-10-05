package models

import "time"

type Channels struct {
	ID          string `json:"id"`
	ChannelType string `json:"channel_type"`
	OwnerID     int    `json:"owner_id"`
	ChannelName string `json:"channel_name"`
	Avatar      string `json:"avatar_url"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
