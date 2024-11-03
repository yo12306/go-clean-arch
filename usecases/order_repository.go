package usecases

import (
	"github.com/go-clean-arch/entities"
)

type OrderRepository interface {
	Save(order entities.Order) error
}
