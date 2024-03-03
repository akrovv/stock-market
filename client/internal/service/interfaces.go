package service

import "github.com/akrovv/client/internal/domain"

type SessionStorage interface {
	Create(dto *domain.Profile) (*domain.Session, error)
	Get(sessionID string) (*domain.Profile, error)
}
