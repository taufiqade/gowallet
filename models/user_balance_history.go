package models

import "time"

// UserBalanceHistory struct
type UserBalanceHistory struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	UserBalanceID uint       `gorm:"column:user_balance_id" json:"user_balance_id"`
	BalanceBefore float64    `gorm:"column:balance_before" gorm:"type:decimal(10,2)" json:"balance_before"`
	BalanceAfter  float64    `gorm:"column:balance_after" gorm:"type:decimal(10,2)"json:"balance_after"`
	Activity      string     `gorm:"column:activity" json:"activity"`
	Type          string     `gorm:"column:type" json:"type"`
	Location      string     `gorm:"column:location" json:"location"`
	IP            string     `gorm:"column:ip" json:"ip"`
	UserAgent     string     `gorm:"column:user_agent" json:"user_agent"`
	Author        string     `gorm:"column:author" json:"author"`
	CreatedAt     *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// IUserBalanceHistoryRepository interface
type IUserBalanceHistoryRepository interface {
	GetBalanceID(id int) (UserBalanceHistory, error)
	Create(data *UserBalanceHistory) error
}
