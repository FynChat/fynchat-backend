package repository

import (
	"context"
	"errors"
	"fmt"

	"fynchat/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (h *UserRepository) Create(ctx context.Context, u *models.User) (*models.User, error) {
	query := `
        INSERT INTO users (name, username, bio, email, password, avatar_url)
        VALUES ($1, $2, $3, $4, $5, NULL)
        RETURNING id
    `
	err := h.pool.QueryRow(ctx, query, u.Name, u.Username, u.Bio, u.Email, u.Password).
		Scan(&u.ID)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return u, nil
}

func (h *UserRepository) FindByID(ctx context.Context, id int) (*models.User, error) {
	query := `SELECT name, username, email FROM users WHERE id = $1`
	u := &models.User{}
	err := h.pool.QueryRow(ctx, query, id).Scan(&u.Name, &u.Username, &u.Email)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("find user by id: %w", err)
	}
	return u, nil
}

func (h *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT name, username, avatar_url FROM users WHERE email = $1`
	u := &models.User{}
	err := h.pool.QueryRow(ctx, query, email).Scan(&u.Name, &u.Username, &u.Avatar)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("find user by email: %w", err)
	}
	return u, nil
}

func (h *UserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `SELECT name, username, avatar_url FROM users WHERE username = $1`
	u := &models.User{}
	err := h.pool.QueryRow(ctx, query, username).Scan(&u.Name, &u.Username, &u.Avatar)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("find user by username: %w", err)
	}
	return u, nil
}

func (h *UserRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := h.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}
	return nil
}

func (h *UserRepository) Update(ctx context.Context) error {
	query := `UPDATE users SET name = $1, username = $2, bio = $3, email = $4, avatar_url = $5`
	u := &models.User{}
	_, err := h.pool.Exec(ctx, query, u.ID, u.Name, u.Username, u.Bio, u.Email, u.Avatar)
	if err != nil {
		return fmt.Errorf("update user: %w", err)
	}
	return nil
}

func (h *UserRepository) GetUserChannels(ctx context.Context, userID int) ([]*models.Channels, error) {
	query := `SELECT * FROM channel_members WHERE user_id = $1
			  INNER JOIN channels ON channel_members.channel_id = channels.id`
	rows, err := h.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("get user channels: %w", err)
	}
	defer rows.Close()

	var channels []*models.Channels
	for rows.Next() {
		var c models.Channels
		err := rows.Scan(&c.ID, &c.ChannelName, &c.Avatar, &c.ChannelType, &c.OwnerID, &c.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("get user channels: %w", err)
		}
		channels = append(channels, &c)
	}

	return channels, nil
}
