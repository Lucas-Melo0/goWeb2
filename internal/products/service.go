package products

type Service interface {
	GetAll() ([]Product, error)
	Insert(id int, name string, color string, price int, stock int, code string, isPublicated bool, creationDate string) (Product, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return []Product{}, err
	}
	return ps, nil
}

func NewService(r Repository) service {
	return &service{
		repository: r,
	}
}
