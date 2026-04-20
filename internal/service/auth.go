package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"

	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/repository"
	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/store"
)

type AuthService struct {
	repo       *repository.Repository
	redis      *redis.Client
	sessionTTL time.Duration
}

func NewAuthService(repo *repository.Repository, redis *redis.Client, sessionTTL time.Duration) *AuthService {
	return &AuthService{
		repo:       repo,
		redis:      redis,
		sessionTTL: sessionTTL,
	}
}

type RegisterInput struct {
	RoleID       string   `json:"roleId"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	Password     string   `json:"password"`
	Phone        string   `json:"phone"`
	Specialty    string   `json:"specialty"`
	MachineFocus []string `json:"machineFocus"`
}

type LoginInput struct {
	RoleID    string `json:"roleId"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	TrainerID string `json:"trainerId"`
}

type AuthResult struct {
	Token   string                  `json:"token,omitempty"`
	Status  string                  `json:"status,omitempty"`
	Message string                  `json:"message,omitempty"`
	User    store.AuthenticatedUser `json:"user,omitempty"`
}

func (s *AuthService) Register(ctx context.Context, payload RegisterInput) (AuthResult, error) {
	payload.RoleID = strings.TrimSpace(payload.RoleID)
	payload.Email = strings.ToLower(strings.TrimSpace(payload.Email))

	if payload.RoleID == "trainer" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
		if err != nil {
			return AuthResult{}, err
		}

		application := store.TrainerApplication{
			ID:           fmt.Sprintf("trainer-application-%s", uuid.NewString()[:8]),
			Name:         strings.TrimSpace(payload.Name),
			Email:        payload.Email,
			Phone:        strings.TrimSpace(payload.Phone),
			Specialty:    strings.TrimSpace(payload.Specialty),
			MachineFocus: payload.MachineFocus,
			Status:       "pending",
			PasswordHash: string(passwordHash),
		}
		if err := s.repo.CreateTrainerApplication(ctx, application); err != nil {
			return AuthResult{}, err
		}

		return AuthResult{
			Status:  "pending",
			Message: "trainer application submitted and is waiting for admin approval",
		}, nil
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return AuthResult{}, err
	}

	user := store.User{
		ID:           fmt.Sprintf("user-%s", uuid.NewString()[:8]),
		Name:         strings.TrimSpace(payload.Name),
		Email:        payload.Email,
		Phone:        strings.TrimSpace(payload.Phone),
		RoleID:       "user",
		PasswordHash: string(passwordHash),
	}
	if err := s.repo.CreateUser(ctx, user); err != nil {
		return AuthResult{}, err
	}

	authUser := sanitizeUser(user, nil)
	token, err := s.createSession(ctx, authUser)
	if err != nil {
		return AuthResult{}, err
	}

	return AuthResult{
		Token:   token,
		Status:  "success",
		Message: "user account created",
		User:    authUser,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, payload LoginInput) (AuthResult, error) {
	payload.Email = strings.ToLower(strings.TrimSpace(payload.Email))

	user, err := s.repo.FindUserByEmail(ctx, payload.Email)
	if err != nil {
		return AuthResult{}, err
	}

	if payload.RoleID != "" && user.RoleID != payload.RoleID {
		return AuthResult{}, errors.New("role mismatch")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(payload.Password)); err != nil {
		return AuthResult{}, errors.New("invalid credentials")
	}

	var trainer *store.Trainer
	if user.RoleID == "trainer" {
		resolvedTrainerID := payload.TrainerID
		if user.TrainerID != nil {
			resolvedTrainerID = *user.TrainerID
		}
		if strings.TrimSpace(resolvedTrainerID) == "" {
			return AuthResult{}, errors.New("trainer account requires trainerId")
		}

		trainerValue, err := s.repo.GetTrainerByID(ctx, resolvedTrainerID)
		if err != nil {
			return AuthResult{}, err
		}
		trainer = &trainerValue
	}

	authUser := sanitizeUser(user, trainer)
	token, err := s.createSession(ctx, authUser)
	if err != nil {
		return AuthResult{}, err
	}

	return AuthResult{
		Token: token,
		User:  authUser,
	}, nil
}

func (s *AuthService) SessionUser(ctx context.Context, token string) (store.AuthenticatedUser, error) {
	raw, err := s.redis.Get(ctx, sessionKey(token)).Bytes()
	if err != nil {
		return store.AuthenticatedUser{}, err
	}

	var user store.AuthenticatedUser
	if err := json.Unmarshal(raw, &user); err != nil {
		return store.AuthenticatedUser{}, err
	}

	return user, nil
}

func (s *AuthService) Logout(ctx context.Context, token string) error {
	return s.redis.Del(ctx, sessionKey(token)).Err()
}

func (s *AuthService) createSession(ctx context.Context, user store.AuthenticatedUser) (string, error) {
	token := uuid.NewString()
	payload, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	if err := s.redis.Set(ctx, sessionKey(token), payload, s.sessionTTL).Err(); err != nil {
		return "", err
	}

	return token, nil
}

func sanitizeUser(user store.User, trainer *store.Trainer) store.AuthenticatedUser {
	result := store.AuthenticatedUser{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.RoleID,
		RoleLabel: roleLabel(user.RoleID),
		TrainerID: user.TrainerID,
	}

	if trainer != nil {
		result.Name = trainer.Name
		trainerID := trainer.ID
		result.TrainerID = &trainerID
	}

	return result
}

func roleLabel(roleID string) string {
	switch roleID {
	case "admin":
		return "admin"
	case "trainer":
		return "trainer"
	default:
		return "user"
	}
}

func sessionKey(token string) string {
	return "session:" + token
}
