package email

type Body struct {
	Email string `json:"email"`
	Otp   int    `json:"otp"`
}
