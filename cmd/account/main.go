package main

import (
	delivery "money-manager/internal/delivery/account"
	
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handler := delivery.NewHandler()
	router := delivery.NewRouter(handler)
	
	lambda.Start(router.Handle)
}