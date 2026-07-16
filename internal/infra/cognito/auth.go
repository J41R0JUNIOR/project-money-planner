package cognito

import (
	"context"
	"errors"
	"log"

	"money-manager/internal/domain"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type CognitoActions struct {
	CognitoClient *cognitoidentityprovider.Client
	ctx           context.Context
	clientId      string
}

func NewCognitoActions(cognitoClient *cognitoidentityprovider.Client) *CognitoActions {
	return &CognitoActions{
		CognitoClient: cognitoClient,
	}
}

func (actor CognitoActions) SignUp(userName string, password string, userEmail string) (bool, error) {
	confirmed := false
	output, err := actor.CognitoClient.SignUp(actor.ctx, &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(actor.clientId),
		Password: aws.String(password),
		Username: aws.String(userName),
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(userEmail)},
		},
	})

	if err != nil {
		if perr, ok := errors.AsType[*types.InvalidPasswordException](err); ok {
			log.Println(perr.Message)
		} else {
			log.Printf("Couldn't sign up user %v. Here's why: %v\n", userName, err)
		}
		
	} else {
		confirmed = output.UserConfirmed
	}

	return confirmed, err
}

func (actor CognitoActions) SignIn(userName string, password string) (domain.Auth, error) {
	var authResult *types.AuthenticationResultType
	output, err := actor.CognitoClient.InitiateAuth(actor.ctx, &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       "USER_PASSWORD_AUTH",
		ClientId:       aws.String(actor.clientId),
		AuthParameters: map[string]string{"USERNAME": userName, "PASSWORD": password},
	})

	if err != nil {

		if perr, ok := errors.AsType[*types.PasswordResetRequiredException](err); ok {
			log.Println(perr.Message)
		} else {
			log.Printf("Couldn't sign in user %v. Here's why: %v\n", userName, err)
		}

	} else {
		authResult = output.AuthenticationResult
	}

	return domain.Auth{
		AccessToken:  *authResult.AccessToken,
		IdToken:      *authResult.IdToken,
		RefreshToken: *authResult.RefreshToken,
	}, err
}