package services

import (
	"context"
	"errors"
	"fmt"
	"fynchat/internal/models"
	"fynchat/internal/repository"
)

var (
	ErrChannelNameAlreadyExists = errors.New("channel name already exists")
)

type ChannelService struct {
	repo *repository.ChannelRepository
}

func NewChannelService(repo *repository.ChannelRepository) *ChannelService {
	return &ChannelService{repo: repo}
}

func (s *ChannelService) Create(ctx context.Context, ownerID int, cType, name, avatar string) (*models.Channels, error) {
	nameExists, err := s.repo.FindByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("check email: %w", err)
	}
	if nameExists != nil {
		return nil, ErrChannelNameAlreadyExists
	}

	channel := &models.Channels{
		ChannelType: cType,
		OwnerID:     ownerID,
		ChannelName: name,
		Avatar:      avatar,
	}
	return s.repo.Create(ctx, channel)
}

func (s *ChannelService) GetMembers(ctx context.Context, channelId string) ([]*models.User, error) {
	members, err := s.repo.GetChannelMembers(ctx, channelId)
	if err != nil {
		return nil, fmt.Errorf("get members: %w", err)
	}
	return members, nil
}

func (s *ChannelService) GetMessages(ctx context.Context, channelId string) ([]*models.Messages, error) {
	messages, err := s.repo.GetChannelMessages(ctx, channelId)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	return messages, nil
}
