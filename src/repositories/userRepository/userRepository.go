package userRepository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/guilhermesants/AppSocialApi/src/entities"
	"github.com/jinzhu/gorm"
)

type Store struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func (svc *Store) Get(id uuid.UUID) (*entities.User, error) {

	var user entities.User
	if err := svc.db.Raw(`SELECT U.* FROM users WHERE id = ?`, id).Scan(&user).Error; err != nil {
		return nil, err
	}

	if user.Id == uuid.Nil {
		return nil, errors.New("Usuario não encontrado.")
	}

	return &user, nil
}

func (svc *Store) Create(user entities.User) (*uuid.UUID, error) {
	return nil, nil
}
