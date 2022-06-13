package car

import "carProject/entity"

type Writer interface {
	Create(c *entity.Car) (entity.ID, error)
}

type Reader interface {
	Get(id entity.ID) (*entity.Car, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateCar(brand string, model string, doorQuantity int) (entity.ID, error)
	GetCar(id entity.ID) (*entity.Car, error)
	//ListCars() ([]*entity.Car, error)
	//SearchCars(query string) ([]*entity.Car, error)
	//UpdateCar(c *entity.Car) error
	//DeleteCar(id entity.ID) error
}
