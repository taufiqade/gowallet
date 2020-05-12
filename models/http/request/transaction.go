package request

// TransactionRequest godoc
type TransactionRequest struct {
	Email     string `json:"email" binding:"required"`
	Amount    int    `json:"amount" binding:"required"`
	IP        string `json:"ip"`
	Location  string `json:"location"`
	UserAgent string `json:"user_agent"`
	Author    string `json:"author"`
}
