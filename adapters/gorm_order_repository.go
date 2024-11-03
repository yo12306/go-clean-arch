package adapters

import (
	"github.com/go-clean-arch/entities"
	"github.com/go-clean-arch/usecases"
	"gorm.io/gorm"
)

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) usecases.OrderRepository {
	return &GormOrderRepository{db: db}
}

func (r *GormOrderRepository) Save(order entities.Order) error {
	return r.db.Create(&order).Error
}
