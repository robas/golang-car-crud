package car

import (
	"carProject/entity"
	"strings"
)

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

func (s *Service) ListCars() ([]*entity.Car, error) {
	c, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *Service) SearchCars(query string) ([]*entity.Car, error) {
	cars, err := s.repo.Search(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(cars) == 0 {
		return nil, entity.ErrNotFound
	}
	return cars, nil
}

func (s *Service) UpdateCar(c *entity.Car) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) DeleteCar(id entity.ID) error {
	//TODO implement me
	panic("implement me")
}
