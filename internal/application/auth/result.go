package auth

type SignInResult struct {
	AccessToken  string
	IdToken      string
	RefreshToken string
}