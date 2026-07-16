package dto

type SignInRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponseDTO struct {
    AccessToken  string `json:"accessToken"`
    IdToken      string `json:"idToken"`
    RefreshToken string `json:"refreshToken"`
}