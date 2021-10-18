package repository

import (
	"github.com/tttinh/goarsenal/entity"
	"gorm.io/gorm"
)

//var RepoErr = errors.New("unable to handle repository request")

type wagerRepositoryImpl struct {
	db *gorm.DB
}

// NewWagerRepository creates a new instance of Wager Repository
func NewWagerRepository(db *gorm.DB) *wagerRepositoryImpl {
	return &wagerRepositoryImpl{db}
}

// FindWagerByID finds a wager by its id.
func (r *wagerRepositoryImpl) FindWagerByID(wagerID uint32) (wager entity.Wager, err error) {
	err = r.db.First(&wager, wagerID).Error
	return
}

// Save creates a new wager
func (r *wagerRepositoryImpl) Save(wager *entity.Wager) error {
	return r.db.Create(wager).Error
}
