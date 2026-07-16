package provider

import "money-manager/internal/domain"

type AuthProvider interface {
	SignUp(userName string, password string, userEmail string) (bool, error)
	SignIn(email string, password string) (domain.Auth, error)
}