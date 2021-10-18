package repository

import (
	"github.com/tttinh/goarsenal/entity"
)

type WagerRepository interface {
	Save(wager *entity.Wager) error
	FindWagerByID(wagerID uint32) (entity.Wager, error)
	FindAll(page uint32, limit uint32) ([]entity.Wager, error)
}

type PurchaseRepository interface {
	Save(purchase *entity.Purchase) error
}
