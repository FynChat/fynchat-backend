package repository

import (
	"context"
	"errors"
	"fmt"
	"fynchat/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ChannelRepository struct {
	pool *pgxpool.Pool
}

func NewChannelRepository(pool *pgxpool.Pool) *ChannelRepository {
	return &ChannelRepository{pool: pool}
}

func (h *ChannelRepository) Create(ctx context.Context, u *models.Channels) (*models.Channels, error) {
	query := `
		INSERT INTO channels(channel_type, owner_id, channel_name, avatar)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`

	err := h.pool.QueryRow(ctx, query, u.ChannelType, u.OwnerID, u.ChannelName, u.Avatar).
		Scan(&u.ID)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return u, nil
}

func (h *ChannelRepository) FindById(ctx context.Context, id string, u *models.Channels) (*models.Channels, error) {
	query := `
		SELECT channel_type, channel_name, avatar_url FROM channels WHERE id = $1
	`
	err := h.pool.QueryRow(ctx, query, id).Scan(&u.ChannelType, &u.ChannelName, &u.Avatar)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("find channel by id: %w", err)
	}
	return u, nil
}

func (h *ChannelRepository) FindByName(ctx context.Context, name string) (*models.Channels, error) {
	query := `
		SELECT channel_type, channel_name, avatar_url FROM channels WHERE channel_name LIKE '%$1%'
	`
	u := &models.Channels{}
	err := h.pool.QueryRow(ctx, query, name).Scan(&u.ChannelType, &u.ChannelName, &u.Avatar)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("find channel by name: %w", err)
	}

	return u, nil
}

func (h *ChannelRepository) GetChannelMembers(ctx context.Context, channelID string) ([]*models.User, error) {
	query := `
		SELECT u.id, u.name, u.username, u.avatar_url
		FROM users u
		JOIN channel_members cm ON u.id = cm.user_id
		WHERE cm.channel_id = $1
	`

	rows, err := h.pool.Query(ctx, query, channelID)
	if err != nil {
		return nil, fmt.Errorf("get channel members: %w", err)
	}
	defer rows.Close()

	var members []*models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Name, &u.Username, &u.Avatar)
		if err != nil {
			return nil, fmt.Errorf("get channel members: %w", err)
		}
		members = append(members, &u)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get channel members: %w", err)
	}

	return members, nil
}

func (h *ChannelRepository) GetChannelMessages(ctx context.Context, channelID string) ([]*models.Messages, error) {
	query := `
		SELECT * FROM messages m INNER JOIN channels c ON m.channel_id = c.id
		WHERE c.id = $1 
	`

	rows, err := h.pool.Query(ctx, query, channelID)
	if err != nil {
		return nil, fmt.Errorf("get channel messages: %w", err)
	}
	defer rows.Close()

	var messages []*models.Messages
	for rows.Next() {
		var m models.Messages
		err := rows.Scan(&m.ID, &m.SenderID, &m.Content, &m.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("get channel messages: %w", err)
		}
		messages = append(messages, &m)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get channel messages: %w", err)
	}
	return messages, nil
}

func (h *ChannelRepository) UpdateChannel(ctx context.Context) error {
	query := `UPDATE channels SET channel_name = $1, avatar_url = $2, updated_at = NOW()`
	c := &models.Channels{}
	_, err := h.pool.Exec(ctx, query, c.ID, c.ChannelName, c.Avatar)
	if err != nil {
		return fmt.Errorf("update channel: %w", err)
	}
	return nil
}

func (h *ChannelRepository) DeleteChannel(ctx context.Context) error {
	query := `DELETE FROM channels WHERE id = $1`
	c := &models.Channels{}
	_, err := h.pool.Exec(ctx, query, c.ID)
	if err != nil {
		return fmt.Errorf("delete channel: %w", err)
	}
	return nil
}
