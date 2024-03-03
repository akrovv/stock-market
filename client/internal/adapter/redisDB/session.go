package redisdb

import (
	"context"
	"encoding/json"
	"time"

	"github.com/akrovv/client/internal/domain"
	"github.com/akrovv/client/pkg/hasher"
	"github.com/redis/go-redis/v9"
)

type sessionStorage struct {
	ctx     context.Context
	storage *redis.Client
	hasher  *hasher.Hasher
}

func NewSessionStorage(ctx context.Context, storage *redis.Client, hasher *hasher.Hasher) *sessionStorage {
	return &sessionStorage{ctx: ctx, storage: storage, hasher: hasher}
}

func (s *sessionStorage) Create(dto *domain.Profile) (*domain.Session, error) {
	emailHash, err := s.hasher.GetHash(dto.Email)

	if err != nil {
		return nil, err
	}

	p, err := json.Marshal(*dto)

	if err != nil {
		return nil, err
	}

	status := s.storage.Set(s.ctx, emailHash, p, time.Hour*5)

	if status.Err() != nil {
		return nil, status.Err()
	}

	return &domain.Session{ID: emailHash}, nil
}

func (s *sessionStorage) Get(sessionID string) (*domain.Profile, error) {
	userProfile := domain.Profile{}

	status := s.storage.Get(s.ctx, sessionID)

	err := status.Err()
	if err != nil {
		return nil, err
	}

	value, err := status.Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(value), &userProfile)

	if err != nil {
		return nil, err
	}

	return &userProfile, nil
}
