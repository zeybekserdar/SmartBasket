package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB,error){
	return gorm.Open(postgres.Open("postgres://user:pass@localhost:5432/SmartBasket"),&gorm.Config{})
}