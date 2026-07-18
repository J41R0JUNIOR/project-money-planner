package dto

type RefreshRequestDTO struct {
	RefreshToken string `json:"refresh_token"`
}

// type RefreshResponseDTO struct {
// 	AccessToken  string `json:"access_token"`
// 	IdToken	  string `json:"id_token"`
// 	RefreshToken string `json:"refresh_token"`
// }