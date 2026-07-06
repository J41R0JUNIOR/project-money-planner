package aws

import "github.com/aws/aws-sdk-go-v2/service/lambda"

type LambdaService struct {
	client *lambda.Client
}

func (f *ServiceFactory) NewLambdaService() *LambdaService {
	return &LambdaService{client: f.NewLambdaClient()}
}
