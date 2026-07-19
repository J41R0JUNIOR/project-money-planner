package dto 

type ConfirmCodeRequestDTO struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}