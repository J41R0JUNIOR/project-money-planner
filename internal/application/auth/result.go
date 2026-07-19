package auth

type AuthResult struct {
	AccessToken  string
	IdToken      string
	RefreshToken string
}
