package db

import (
	"github.com/jinzhu/gorm"
	"github.com/taufiqade/gowallet/models"
)

// UserBalanceRepository godoc
type UserBalanceRepository struct {
	DB *gorm.DB
}

// NewUserBalanceRepository godoc
func NewUserBalanceRepository(db *gorm.DB) models.IUserBalanceRepository {
	return &UserBalanceRepository{DB: db}
}

// GetByUserID godoc
func (repo *UserBalanceRepository) GetByUserID(userID int) (models.UserBalance, error) {
	var ub models.UserBalance
	query := repo.DB.Table("user_balance").Where("user_id=?", userID).First(&ub)
	return ub, query.Error
}

// Update godoc
func (repo *UserBalanceRepository) Update(userID int, data *models.UserBalance) error {
	//var ub models.UserBalance
	query := repo.DB.Table("user_balance").Where("user_id=?", userID).Updates(data)
	return query.Error
}

// Create godoc
func (repo *UserBalanceRepository) Create(data *models.UserBalance, tx *gorm.DB) error {
	query := tx.Table("user_balance").Create(data)
	return query.Error
}
