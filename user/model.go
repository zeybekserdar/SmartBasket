package user

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         uuid.UUID `json:"id" gorm:"primary_key;uuid"`
	FirstName  string    `json:"firstName" validate:"required"`
	LastName   string    `json:"lastName" validate:"required"`
	Username   string    `json:"username" validate:"required"`
	Email      string    `json:"email" validate:"email,required"`
	Password   string    `json:"password" validate:"required"`
	gorm.Model `json:"-"`
}

func (user *User) BeforeCreate(*gorm.DB) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}
