package repository

import (
	"context"
	"fmt"
	"fynchat.thegrigoriyagen.com/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MessagesRepository struct {
	pool *pgxpool.Pool
}

func NewMessagesRepository(pool *pgxpool.Pool) *MessagesRepository {
	return &MessagesRepository{pool: pool}
}

func (r *MessagesRepository) Create(ctx context.Context, userId int, content string, channelId uuid.UUID) (*models.Messages, error) {
	query := `INSERT INTO messages (channel_id, sender_id, content, created_at)
			  VALUES($1, $2, $3, NOW())`
	c := &models.Messages{}
	err := r.pool.QueryRow(ctx, query, c.ChannelID, c.SenderID, c.Content).
		Scan(&c.ID)
	if err != nil {
		return nil, fmt.Errorf("create message: %w", err)
	}
	return c, nil
}

func (r *MessagesRepository) Delete(ctx context.Context, MessageID string) error {
	query := `DELETE FROM messages WHERE id = $1`
	_, err := r.pool.Exec(ctx, query, MessageID)
	if err != nil {
		return fmt.Errorf("delete message: %w", err)
	}
	return nil
}

func (r *MessagesRepository) IsMember(ctx context.Context, senderID int64, channelID uuid.UUID) (bool, error) {
	query := `SELECT EXISTS (
    SELECT 1 FROM channel_members
    WHERE channel_id = $1 AND user_id = $2
)`
	var exists bool
	err := r.pool.QueryRow(ctx, query, channelID, senderID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check is member: %w", err)
	}
	return exists, nil
}
