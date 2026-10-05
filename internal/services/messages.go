package services

import (
	"context"
	"errors"
	"fmt"
	"fynchat/internal/models"
	"fynchat/internal/repository"

	"github.com/google/uuid"
)

var (
	ErrNotMember = errors.New("Member don't exists")
)

type MessageService struct {
	repo *repository.MessagesRepository
}

func NewMessageService(repo *repository.MessagesRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (m *MessageService) Create(ctx context.Context, senderID int64, channelID uuid.UUID, content string) (*models.Messages, error) {
	IsMember, err := m.repo.IsMember(ctx, senderID, channelID)
	if err != nil {
		return nil, fmt.Errorf("check membership: %w", err)
	}

	if !IsMember {
		return nil, ErrNotMember
	}

	message, err := m.repo.Create(ctx, int(senderID), content, channelID)
	if err != nil {
		return nil, fmt.Errorf("create message: %w", err)
	}
	return message, nil
}
