package car

import (
	"carProject/entity"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func newFakeCar() *entity.Car {
	return &entity.Car{
		Brand:        "Volkswagen",
		Model:        "Fusca",
		DoorQuantity: 4,
		CreatedAt:    time.Now(),
	}
}

func TestService_CreateCar(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	c := newFakeCar()
	_, err := s.CreateCar(c.Brand, c.Model, c.DoorQuantity)
	assert.Nil(t, err)
	assert.False(t, c.CreatedAt.IsZero())
}

func TestService_GetCar(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	c1 := newFakeCar()
	c2 := newFakeCar()
	cID, _ := s.CreateCar(c1.Brand, c1.Model, c1.DoorQuantity)
	_, _ = s.CreateCar(c2.Brand, c2.Model, c2.DoorQuantity)
	saved, err := s.GetCar(cID)
	assert.Nil(t, err)
	assert.Equal(t, c1.Brand, saved.Brand)
	assert.Equal(t, c1.Model, saved.Model)
	assert.Equal(t, c1.DoorQuantity, saved.DoorQuantity)
}
