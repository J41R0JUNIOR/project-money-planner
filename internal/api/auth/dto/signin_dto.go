package auth_dto

type SignInRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponseDTO struct {
	Message string `json:"message"`
}