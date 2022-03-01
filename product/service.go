package product

type Service interface {
	Get(id uint) (*Product,error)
	Create(model Product) (uint,error)
}

type service struct {
	repository Repository
}

func (s service) Get(id uint) (*Product, error) {
	return s.repository.Get(id)
}

func (s service) Create(product Product) (uint, error) {
	return s.repository.Create(product)
}

var _ Service = service{}

func NewService(repository Repository) Service{
	return service{repository: repository}
}