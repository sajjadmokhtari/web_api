package dto

// GetOtpRequest is used for requesting OTP
// swagger:model GetOtpRequest
//
type GetOtpRequest struct {
	// Mobile number of the user
	MobileNumber string `json:"mobileNumber" binding:"required,min=11,max=11"`
}

// TokenDetail contains access and refresh tokens
// swagger:model TokenDetail
//
type TokenDetail struct {
	AccessToken            string `json:"accessToken"`
	RefreshToken           string `json:"refreshToken"`
	AccessTokenExpireTime  int64  `json:"accessTokenExpireTime"`
	RefreshTokenExpireTime int64  `json:"refreshTokenExpireTime"`
}

// RegisterUserByUsernameRequest is used for registering a user by username
// swagger:model RegisterUserByUsernameRequest
//
type RegisterUserByUsernameRequest struct {
	FirstName string `json:"firstName" binding:"required,min=3"`
	LastName  string `json:"lastName"  binding:"required,min=6"`
	Username  string `json:"username"  binding:"required,min=5"`
	Email     string `json:"email" binding:"min=6,email"`
	Password  string `json:"password" binding:"required,password,min=6"`
}

// RegisterLoginByMobileRequest is used for registering or logging in by mobile
// swagger:model RegisterLoginByMobileRequest
//
type RegisterLoginByMobileRequest struct {
	MobileNumber string `json:"mobileNumber" binding:"required,mobile,min=11,max=11"`
	Otp          string `json:"otp" binding:"required,min=6,max=6"`
}

// LoginByUsernameRequest is used for logging in by username
// swagger:model LoginByUsernameRequest
//
type LoginByUsernameRequest struct {
	Username string `json:"username" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=6"`
}
