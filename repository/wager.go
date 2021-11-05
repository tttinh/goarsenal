package repository

import (
	"github.com/tttinh/goarsenal/entity"
	"gorm.io/gorm"
)

// var RepoErr = errors.New("unable to handle repository request")

type wagerRepositoryImpl struct {
	db *gorm.DB
}

// NewWagerRepository creates a new instance of Wager Repository.
func NewWagerRepository(db *gorm.DB) *wagerRepositoryImpl {
	return &wagerRepositoryImpl{db}
}

// Create creates a new wager.
func (r *wagerRepositoryImpl) Create(wager *entity.Wager) error {
	return r.db.Create(wager).Error
}

// Update update an existing wager.
func (r *wagerRepositoryImpl) Update(wager *entity.Wager) error {
	return r.db.Save(wager).Error
}

// FindWagerByID finds a wager by its id.
func (r *wagerRepositoryImpl) FindWagerByID(wagerID uint32) (wager entity.Wager, err error) {
	err = r.db.First(&wager, wagerID).Error
	return
}

// FindAll finds all wagers with page and limit.
func (r *wagerRepositoryImpl) FindAll(page uint32, limit uint32) (wagers []entity.Wager, err error) {
	err = r.db.Limit(int(limit)).Offset(int(page * limit)).Find(&wagers).Error
	return
}
