package user

import (
	"github.com/gofrs/uuid"
)

type IUserService interface {
	Get(id uuid.UUID) (*User, error)
	GetByUsername(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(model User) (*User, error)
	Update(old User, new User) (uuid.UUID, error)
	Delete(model User) (uuid.UUID, error)
}

type service struct {
	repository UserRepository
}

func (s service) GetByUsername(username string) (*User, error) {
	return s.repository.GetByUsername(username)
}

func (s service) GetByEmail(email string) (*User, error) {
	return s.repository.GetByEmail(email)
}
func (s service) Update(old User, new User) (uuid.UUID, error) {
	return s.repository.Update(old, new)
}

func (s service) Delete(model User) (uuid.UUID, error) {
	return s.repository.Delete(model)
}

func (s service) Get(id uuid.UUID) (*User, error) {
	return s.repository.Get(id)
}

func (s service) Create(product User) (*User, error) {
	return s.repository.Create(product)
}

var _ IUserService = service{}

func NewService(repository UserRepository) IUserService {
	return service{repository: repository}
}
