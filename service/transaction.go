package service

import (
	dbEntity "github.com/taufiqade/gowallet/models"
	"github.com/taufiqade/gowallet/models/http/request"
)

type transactionService struct {
	userRepo 			dbEntity.IUserRepository
	userBalanceRepo 	dbEntity.IUserBalanceRepository
	historyRepo			dbEntity.IUserBalanceHistoryRepository
}

func NewTransactionService(u dbEntity.IUserRepository, ub dbEntity.IUserBalanceRepository, uh dbEntity.IUserBalanceHistoryRepository) *transactionService {
	return &transactionService{
		userRepo: u,
		userBalanceRepo: ub,
		historyRepo: uh,
	}
}

func (u *transactionService) TopUp(userId int, payload *request.TopUpRequest) error {
	beneficiary, err:= u.userBalanceRepo.GetByUserID(userId)
	if err != nil {
		return err
	}
	currBalance := beneficiary.Balance
	// update balanceData
	balanceData:= &dbEntity.UserBalance{
		ID:              	beneficiary.ID,
		UserID:          	beneficiary.UserID,
		Balance:         	beneficiary.Balance + float64(payload.Amount),
		BalanceAchieve: 	beneficiary.BalanceAchieve + float64(payload.Amount),
	}
	ubErr := u.userBalanceRepo.Update(userId, balanceData)
	if ubErr != nil {
		return ubErr
	}
	//insert balance history
	balanceHistory:= &dbEntity.UserBalanceHistory{
		UserBalanceID: 	beneficiary.ID,
		BalanceBefore: 	currBalance,
		BalanceAfter: 	balanceData.BalanceAchieve,
		Activity:      	"TopUp",
		Type:  			"debit",
		IP:            	payload.IP,
		UserAgent:     	payload.UserAgent,
		Location: 		payload.Location,
		Author: 		payload.Author,
	}
	bhErr := u.historyRepo.Create(balanceHistory)
	if bhErr != nil {
		return bhErr
	}
	return bhErr
}