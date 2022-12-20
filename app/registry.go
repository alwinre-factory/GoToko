package app

import "github.com/alwinre-factory/GoToko/app/models"

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Address{}},
		{Model: models.Product{}},
		{Model: models.Payment{}},
		{Model: models.ProductImage{}},
		{Model: models.Category{}},
		{Model: models.OrderCustomer{}},
		{Model: models.OrderItem{}},
		{Model: models.Order{}},
		{Model: models.Section{}},
		{Model: models.Shipment{}},
	}
}
