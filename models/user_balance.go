package models

import "github.com/jinzhu/gorm"

// UserBalance struct
type UserBalance struct {
	ID             uint    `gorm:"primary_key" json:"id"`
	UserID         uint    `gorm:"column:user_id" json:"user_id"`
	Balance        float64 `gorm:"column:balance" gorm:"type:decimal(10,2)" json:"balance"`
	BalanceAchieve float64 `gorm:"column:balance_achieve" gorm:"type:decimal(10,2)" json:"balance_achieve"`
}

// IUserBalanceRepository contract
type IUserBalanceRepository interface {
	GetByUserID(userID int) (UserBalance, error)
	Update(userID int, data *UserBalance) error
	Create(data *UserBalance, tx *gorm.DB) error
}
