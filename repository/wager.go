package repository

import (
	"github.com/tttinh/goarsenal/entity"
	"gorm.io/gorm"
)

type WagerRepository interface {
	AddWager(wager *entity.Wager) error
}

//var RepoErr = errors.New("unable to handle repository request")

type wagerRepositoryImpl struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *wagerRepositoryImpl {
	return &wagerRepositoryImpl{
		db: db,
	}
}

// AddWager adds a new wager
func (r *wagerRepositoryImpl) AddWager(wager *entity.Wager) error {
	if err := r.db.Create(wager).Error; err != nil {
		return err
	}

	return nil
}
