package car

import "carProject/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateCar(brand string, model string, doorQuantity int) (entity.ID, error) {
	c, err := entity.NewCar(brand, model, doorQuantity)
	if err != nil {
		return entity.NewID(), err
	}
	return s.repo.Create(c)
}

func (s *Service) GetCar(id entity.ID) (*entity.Car, error) {
	c, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, entity.ErrNotFound
	}

	return c, nil
}