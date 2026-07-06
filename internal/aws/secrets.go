package aws

import "github.com/aws/aws-sdk-go-v2/service/secretsmanager"

type SecretsService struct {
	client *secretsmanager.Client
}

func (f *ServiceFactory) NewSecretsService() *SecretsService {
	return &SecretsService{client: f.NewSecretsManagerClient()}
}
