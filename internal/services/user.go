package services

import (
	"context"
	"errors"
	"fmt"

	"fynchat.thegrigoriyagen.com/internal/models"
	"fynchat.thegrigoriyagen.com/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrUserNotFound          = errors.New("user not found")
	ErrUsernameAlreadyExists = errors.New("username already exists")
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, name, username, email, password string) (*models.User, error) {
	emailExists, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("check email: %w", err)
	}
	if emailExists != nil {
		return nil, ErrEmailAlreadyExists
	}

	usernameExists, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("check username: %w", err)
	}
	if usernameExists != nil {
		return nil, ErrUsernameAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	user := &models.User{
		Name:     name,
		Username: username,
		Email:    email,
		Password: string(hash),
	}
	return s.repo.Create(ctx, user)
}

func (s *UserService) Login (ctx context.Context, email, password string) (*models.User, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("find user: %w", err)
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

func (s *UserService) GetById(ctx context.Context, id int) (*models.User, error) {
	user, err := s.repo.FindByID(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("find user: %w", err)
	}

	if user == nil {
		 return nil, ErrUserNotFound
	}
	return user, nil
}

func (s *UserService) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := s.repo.FindByUsername(ctx, username)
	
	if err != nil {
		return nil, fmt.Errorf("find user by username: %w", err)
	}

	if user == nil {
		 return nil, ErrUserNotFound
	}

	return user, nil
}

func (s *UserService) Update(ctx context.Context, id int, name, username, email string) (*models.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find user: %w", err)
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	user.Name = name
	user.Username = username
	user.Email = email

	if err := s.repo.Update(ctx); err != nil {
		return nil, fmt.Errorf("update user: %w", err)
	}
	return user, nil
}

func (s *UserService) Delete(ctx context.Context, id int) error {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	if user == nil {
		return ErrUserNotFound
	}

	return s.repo.Delete(ctx, id)
}
