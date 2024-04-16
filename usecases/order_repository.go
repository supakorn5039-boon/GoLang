package usecases

import "boon/entities"

type OrderRepository interface {
	Save(order entities.Order) error
}
