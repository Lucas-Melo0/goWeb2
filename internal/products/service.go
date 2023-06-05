package products

type Service interface {
	GetAll() ([]Product, error)
	Insert(name string, color string, price int, stock int, code string, isPublicated bool, creationDate string) (Product, error)
	Update(id int, name string, color string, price int, stock int, code string, isPublicated bool, creationDate string) (Product, error)
	Delete(id int) error
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
func (s *service) Insert(name string, color string, price int, stock int, code string, isPublicated bool, creationDate string) (Product, error) {
	id, err := s.repository.LastId()
	if err != nil {
		return Product{}, err
	}
	id++
	p, err := s.repository.Insert(id, name, color, price, stock, code, isPublicated, creationDate)

	return p, err
}
func (s *service) Update(id int, name string, color string, price int, stock int, code string, isPublicated bool, creationDate string) (Product, error) {
	p, err := s.repository.Update(id, name, color, price, stock, code, isPublicated, creationDate)
	return p, err
}
func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
