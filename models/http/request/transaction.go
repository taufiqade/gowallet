package request

type TopUpRequest struct {
	BeneficiaryId int `json:"beneficiary_id" binding:"required"`
	Amount int `json:"amount" binding:"required"`
	IP string `json:"ip"`
	Location string `json:"location"`
	UserAgent string `json:"user_agent"`
	Author string `json:"author"`
}

//type TransferRequest struct {
//	BeneficiaryId int `json:"beneficiary_id"`
//	Amount        int `json:"amount"`
//}