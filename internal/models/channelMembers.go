package models

type ChannelMembers struct {
	ID        string `json:"id"`
	ChannelID string `json:"channel_id"`
	UserID    int    `json:"user_id"`
}
