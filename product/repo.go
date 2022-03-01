package product

import "gorm.io/gorm"

type Repository interface{
	Get(id uint) (*Product,error)
	Create(model Product) (uint,error)
	Migration() error
}

type repository struct {
	db *gorm.DB
}

func (r repository) Get(id uint) (*Product, error) {
	product := &Product{ID: id}
	err := r.db.First(product).Error
	if err != nil {
		return nil,err
	}
	return product,nil
}

func (r repository) Create(product Product) (uint, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return 0,err
	}
	return product.ID,nil
}

func (r repository) Migration() error {
	return r.db.AutoMigrate(&Product{})
}

var _ Repository = repository{}

func NewRepository(db *gorm.DB) Repository{
return repository{db:db}
}

