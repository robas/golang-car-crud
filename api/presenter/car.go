package presenter

import (
	"carProject/entity"
)

type Car struct {
	ID           entity.ID `json:"id"`
	Brand        string    `json:"brand"`
	Model        string    `json:"model"`
	DoorQuantity int       `json:"doorQuantity"`
}
