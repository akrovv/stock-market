package rest

import (
	"github.com/akrovv/client/internal/domain"
	"github.com/akrovv/client/internal/service"
)

type UserService interface {
	Save(dto *service.SaveUser) error
	Get(dto *service.GetUser) (*domain.Profile, error)
}

type SessionService interface {
	Create(dto *domain.Profile) (*domain.Session, error)
	Get(dto *service.GetSession) (*domain.Profile, error)
}
