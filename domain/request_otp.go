package domain

type RequestOTP struct {
	Email string `json:"email" binding:"required" example:"user4@mailinator.com"`
}
