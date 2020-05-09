package db

import (
	"github.com/jinzhu/gorm"
	"test/models"
)

type UserBalanceRepository struct {
	DB *gorm.DB
}

func NewUserBalanceRepository(db *gorm.DB) models.IUserBalanceRepository {
	return &UserBalanceRepository{DB: db}
}

func (repo *UserBalanceRepository) GetByUserID(userId int) (models.UserBalance, error) {
	var ub models.UserBalance
	query := repo.DB.Table("user_balance").Where("user_id=?", userId).First(&ub)
	return ub, query.Error
}

func (repo *UserBalanceRepository) Update(userId int, data *models.UserBalance) error {
	//var ub models.UserBalance
	query := repo.DB.Table("user_balance").Where("user_id=?",userId).Updates(data)
	return query.Error
}