package dto

type SignUpRequestDTO struct {
	Email    string `json:"email"`
	Name	 string `json:"name"`
	Password string `json:"password"`
}

// type SignUpResponseDTO struct {
// 	Message string `json:"message"`
// }