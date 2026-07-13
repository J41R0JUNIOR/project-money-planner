package auth_dto

type SignUpRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponseDTO struct {
	Message string `json:"message"`
}