package dto


type GetOtpRequest struct {
    MobileNumber string `json:"mobileNumber" binding:"required,min=11,max=11"`
}
