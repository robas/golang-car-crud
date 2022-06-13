package entity

import "time"

type Car struct {
	ID           ID
	Brand        string
	Model        string
	DoorQuantity int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewCar(brand string, model string, doorQuantity int) (*Car, error) {
	c := &Car{
		ID:           NewID(),
		Brand:        brand,
		Model:        model,
		DoorQuantity: doorQuantity,
		CreatedAt:    time.Now(),
	}
	err := c.Validate()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Car) Validate() error {
	if c.DoorQuantity <= 0 || c.Brand == "" || c.Model == "" {
		return ErrInvalidEntity
	}
	return nil
}
