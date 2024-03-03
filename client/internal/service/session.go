package service

import (
	"github.com/akrovv/client/internal/domain"
)

type sessionService struct {
	storage SessionStorage
}

func NewSessionService(storage SessionStorage) *sessionService {
	return &sessionService{storage: storage}
}

func (s *sessionService) Create(dto *domain.Profile) (*domain.Session, error) {
	return s.storage.Create(dto)
}

func (s *sessionService) Get(dto *GetSession) (*domain.Profile, error) {
	return s.storage.Get(dto.ID)
}
