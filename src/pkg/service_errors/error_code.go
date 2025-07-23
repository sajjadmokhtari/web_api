package service_errors

const(
    //TOKEN
	UnExpectedError = "Expected error"
	ClaimsNotFound = "Claims Not Found"
	TokenExpired = "Token Expired"
	TokenRequired = "Token Required"
	TokenInvalid = "Token Invalid"


	// OTP
	OtpExists = "Otp exists"
	OtpUsed = "Otp Used"
	OtpNotValid = "Otp not Valid"

	//user

	EmailExists = "Email exists"
	UsernameExists = "UsernameExists"
	PermissionDenied = "Permission Denied"
	
	//Db

	RecordNotFound = "record not found"
	
) 