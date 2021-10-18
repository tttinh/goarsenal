package repository

import (
	"github.com/tttinh/goarsenal/entity"
	"gorm.io/gorm"
)

type purchaseRepositoryImpl struct {
	db *gorm.DB
}

// NewPurchaseRepository creates a new instance of Purchase repository
func NewPurchaseRepository(db *gorm.DB) *purchaseRepositoryImpl {
	return &purchaseRepositoryImpl{db}
}

// Save creates a new purchase
func (r *purchaseRepositoryImpl) Save(purchase *entity.Purchase) error {
	return r.db.Create(purchase).Error
}
