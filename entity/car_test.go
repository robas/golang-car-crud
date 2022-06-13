package entity_test

import (
	"carProject/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCar(t *testing.T) {
	brand := "Volkswagen"
	model := "Fusca"

	c, err := entity.NewCar(brand, model, 4)
	assert.Nil(t, err)
	assert.Equal(t, c.Brand, brand)
	assert.NotNil(t, c.ID)
}

func TestCar_Validate(t *testing.T) {
	type test struct {
		brand        string
		model        string
		doorQuantity int
		want         error
	}

	tests := []test{
		{
			brand:        "",
			model:        "Fusca",
			doorQuantity: 4,
			want:         entity.ErrInvalidEntity,
		},
		{
			brand:        "Volkswagen",
			model:        "",
			doorQuantity: 4,
			want:         entity.ErrInvalidEntity,
		},
		{
			brand:        "Volkswagen",
			model:        "Fusca",
			doorQuantity: 0,
			want:         entity.ErrInvalidEntity,
		},
		{
			brand:        "Volkswagen",
			model:        "Fusca",
			doorQuantity: -1,
			want:         entity.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {
		_, err := entity.NewCar(tc.brand, tc.model, tc.doorQuantity)
		assert.Equal(t, err, tc.want)
	}
}
