package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	Get(id uuid.UUID) (*User, error)
	Create(user User) (*uuid.UUID, error)
}

type User struct {
	Id                uuid.UUID `gorm:"column:id"`
	Login             string    `gorm:"column:login"`
	Password          string    `gorm:"column:password"`
	Email             string    `gorm:"column:email"`
	Name              string    `gorm:"column:name"`
	LastName          string    `gorm:"column:lastName"`
	ProfileImage      []byte    `gorm:"column:profileImage"`
	LastLogin         time.Time `gorm:"column:lastLogin"`
	ConfirmDataAccess bool      `gorm:"column:confirm_data_access"`
}

func (User) TableName() string {
	return "users"
}
