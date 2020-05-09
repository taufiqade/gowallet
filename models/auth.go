package models

// TokenDetails is represent AuthDetail
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

// IAuthService is transaction service contract
type IAuthService interface {
	Login(email string, password string) (string, error)
}
