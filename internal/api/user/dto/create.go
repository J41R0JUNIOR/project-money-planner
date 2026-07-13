package dto

type CreateUserRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponseDTO struct {
	Message string `json:"message"`
}