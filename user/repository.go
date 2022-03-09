package user

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Get(id uuid.UUID) (*User, error)
	GetByUsername(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(model User) (*User, error)
	Update(old User, new User) (uuid.UUID, error)
	Delete(model User) (uuid.UUID, error)
	Migration() error
}

type userRepository struct {
	db *gorm.DB
}

func (r userRepository) GetByUsername(username string) (*User, error) {
	user := &User{}
	err := r.db.Where(&User{Username: username}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r userRepository) GetByEmail(email string) (*User, error) {
	user := &User{}
	err := r.db.Where(&User{Email: email}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r userRepository) Update(old User, new User) (uuid.UUID, error) {
	err := r.db.Model(old).Updates(new).Error
	if err != nil {
		return uuid.Nil, err
	}
	return old.ID, nil
}

func (r userRepository) Delete(model User) (uuid.UUID, error) {
	err := r.db.Delete(&model).Error
	if err != nil {
		return uuid.Nil, err
	}
	return model.ID, nil
}

func (r userRepository) Get(id uuid.UUID) (*User, error) {
	user := &User{ID: id}
	err := r.db.First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r userRepository) Create(model User) (*User, error) {
	err := r.db.Create(&model).Error
	if err != nil {
		return &User{}, err
	}
	return &model, nil
}

func (r userRepository) Migration() error {
	return r.db.AutoMigrate(&User{})
}

var _ UserRepository = userRepository{}

func NewRepository(db *gorm.DB) UserRepository {
	return userRepository{db: db}
}
