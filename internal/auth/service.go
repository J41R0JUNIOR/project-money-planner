package auth

// Service expõe a lógica de autenticação inicial do módulo.
type Service struct{}

// NewService cria o serviço de autenticação.
func NewService() *Service {
	return &Service{}
}
