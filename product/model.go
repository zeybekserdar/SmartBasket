package product

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID         uuid.UUID `json:"id" gorm:"primary_key;uuid"`
	Name       string    `json:"name" validate:"required"`
	Price      string    `json:"price"`
	Type       string    `json:"type"`
	gorm.Model `json:"-"`
}

func (product *Product) BeforeCreate() error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	product.ID = id
	return nil
}
