package repository

import (
	"github.com/tttinh/goarsenal/entity"
)

type WagerRepository interface {
	Update(wager *entity.Wager) error
	Create(wager *entity.Wager) error
	FindWagerByID(wagerID uint32) (entity.Wager, error)
	FindAll(page uint32, limit uint32) ([]entity.Wager, error)
}

type PurchaseRepository interface {
	Create(purchase *entity.Purchase) error
}
