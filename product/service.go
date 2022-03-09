package product

import (
	"github.com/gofrs/uuid"
)

type IProductService interface {
	Get(id uuid.UUID) (*Product, error)
	GetAll() (*[]Product, error)
	Create(model Product) (*Product, error)
	Update(old Product, new Product) (uuid.UUID, error)
	Delete(model Product) (uuid.UUID, error)
}

type service struct {
	repository Repository
}

func (s service) GetAll() (*[]Product, error) {
	return s.repository.GetAll()
}

func (s service) Update(old Product, new Product) (uuid.UUID, error) {
	return s.repository.Update(old, new)
}

func (s service) Delete(model Product) (uuid.UUID, error) {
	return s.repository.Delete(model)
}

func (s service) Get(id uuid.UUID) (*Product, error) {
	return s.repository.Get(id)
}

func (s service) Create(product Product) (*Product, error) {
	return s.repository.Create(product)
}

var _ IProductService = service{}

func NewService(repository Repository) IProductService {
	return service{repository: repository}
}
