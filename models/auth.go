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

// AccessDetails godoc
type AccessDetails struct {
	AccessUUID string
	UserID     uint64
}

// IAuthService is transaction service contract
type IAuthService interface {
	CreateToken(email string, password string) (string, error)
}

// IRedisAuthRepository godoc
type IRedisAuthRepository interface {
	Get(key string) (string, error)
	Set(key, value string, exp int64) (err error)
}
