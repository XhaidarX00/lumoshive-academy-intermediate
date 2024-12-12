package models

type OTPData struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type OTPDataKey struct {
	Key string `json:"valid_key"`
	OTP string `json:"otp"`
}
