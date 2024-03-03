package service

import (
	"github.com/akrovv/client/internal/domain"
)

type userService struct {
}

func NewUserService() *userService {
	return &userService{}
}

func (s *userService) Save(dto *SaveUser) error {
	return nil
}

func (s *userService) Get(dto *GetUser) (*domain.Profile, error) {
	return &domain.Profile{Nickname: "Akro", Email: "akro.info@mail.ru", Firstname: "Artyom", Lastname: "Blokhin"}, nil
}
