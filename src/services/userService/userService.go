package userService

import (
	"github.com/google/uuid"
	"github.com/guilhermesants/AppSocialApi/src/entities"
)

type UserService interface {
	Get(id uuid.UUID) (*entities.User, error)
	Create(user entities.User) (*uuid.UUID, error)
}

type Service struct {
	repository entities.UserRepository
}

func NewUserService(repository entities.UserRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (svc *Service) Get(id uuid.UUID) (*entities.User, error) {

	user, err := svc.repository.Get(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (svc *Service) Create(user entities.User) (*uuid.UUID, error) {

	return nil, nil
}
