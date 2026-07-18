package cognito

import (
	"context"
	"errors"
	"log"

	domain "money-manager/internal/domain/auth"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type CognitoActions struct {
	CognitoClient *cognitoidentityprovider.Client
	clientId      string
}

func NewCognitoActions(cognitoClient *cognitoidentityprovider.Client, clientId string) *CognitoActions {
	return &CognitoActions{
		CognitoClient: cognitoClient,
		clientId:      clientId,
	}
}

func (actor CognitoActions) SignUp(ctx context.Context, userName string, password string, userEmail string) (string, error) {
	result, err := actor.CognitoClient.SignUp(ctx, &cognitoidentityprovider.SignUpInput{
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
		
	} 

	return *result.UserSub, err
}

func (actor CognitoActions) SignIn(ctx context.Context, userName string, password string) (domain.Auth, error) {
	var authResult *types.AuthenticationResultType
	output, err := actor.CognitoClient.InitiateAuth(ctx, &cognitoidentityprovider.InitiateAuthInput{
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

func (actor CognitoActions) Refresh(ctx context.Context, refreshToken string) (domain.Auth, error) {
	var authResult *types.AuthenticationResultType
	output, err := actor.CognitoClient.InitiateAuth(ctx, &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       "REFRESH_TOKEN_AUTH",
		ClientId:       aws.String(actor.clientId),
		AuthParameters: map[string]string{"REFRESH_TOKEN": refreshToken},
	})

	if err != nil {
		log.Printf("Couldn't refresh token. Here's why: %v\n", err)
	} else {
		authResult = output.AuthenticationResult
	}

	return domain.Auth{
		AccessToken:  *authResult.AccessToken,
		IdToken:      *authResult.IdToken,
		RefreshToken: refreshToken,
	}, err
}