package product

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Get(id uuid.UUID) (*Product, error)
	GetAll() (*[]Product, error)
	Create(model Product) (*Product, error)
	Update(old Product, new Product) (uuid.UUID, error)
	Delete(model Product) (uuid.UUID, error)
	Migration() error
}

type repository struct {
	db *gorm.DB
}

func (r repository) GetAll() (*[]Product, error) {
	allProducts := []Product{}
	err := r.db.Find(&allProducts).Error
	if err != nil {
		return nil, err
	}
	return &allProducts, nil
}

func (r repository) Update(old Product, new Product) (uuid.UUID, error) {
	err := r.db.Model(old).Updates(new).Error
	if err != nil {
		return uuid.Nil, err
	}
	return old.ID, nil
}

func (r repository) Delete(model Product) (uuid.UUID, error) {
	err := r.db.Delete(&model).Error
	if err != nil {
		return uuid.Nil, err
	}
	return model.ID, nil
}

func (r repository) Get(id uuid.UUID) (*Product, error) {
	product := &Product{ID: id}
	err := r.db.First(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r repository) Create(model Product) (*Product, error) {
	err := r.db.Create(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (r repository) Migration() error {
	return r.db.AutoMigrate(&Product{})
}

var _ Repository = repository{}

func NewRepository(db *gorm.DB) Repository {
	return repository{db: db}
}
