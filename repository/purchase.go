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

// Create creates a new purchase
func (r *purchaseRepositoryImpl) Create(purchase *entity.Purchase) error {
	return r.db.Create(purchase).Error
}
