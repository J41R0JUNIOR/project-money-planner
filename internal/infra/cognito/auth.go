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

func (actor CognitoActions) SignUp(ctx context.Context, userEmail string, password string) (string, error) {
	result, err := actor.CognitoClient.SignUp(ctx, &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(actor.clientId),
		Password: aws.String(password),
		Username: aws.String(userEmail),
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(userEmail),
			},
		},
	})

	if err != nil {
		if perr, ok := errors.AsType[*types.InvalidPasswordException](err); ok {
			log.Println(*perr.Message)
		} else {
			log.Printf("Couldn't sign up user %v. Here's why: %v\n", userEmail, err)
		}

		return "", err
	}

	return aws.ToString(result.UserSub), nil
}

func (actor CognitoActions) ConfirmCode(ctx context.Context, email string, code string) error {
	_, err := actor.CognitoClient.ConfirmSignUp(ctx, &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(actor.clientId),
		Username:         aws.String(email),
		ConfirmationCode: aws.String(code),
	})

	if err != nil {
		log.Printf("Couldn't confirm code for user %v. Here's why: %v\n", email, err)
		return err
	}

	return nil
}

func (actor CognitoActions) SignIn(ctx context.Context, userName string, password string) (domain.Auth, error) {
	output, err := actor.CognitoClient.InitiateAuth(ctx, &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: "USER_PASSWORD_AUTH",
		ClientId: aws.String(actor.clientId),
		AuthParameters: map[string]string{
			"USERNAME": userName,
			"PASSWORD": password,
		},
	})

	if err != nil {
		if perr, ok := errors.AsType[*types.PasswordResetRequiredException](err); ok {
			log.Println(*perr.Message)
		} else {
			log.Printf("Couldn't sign in user %v. Here's why: %v\n", userName, err)
		}

		return domain.Auth{}, err
	}

	authResult := output.AuthenticationResult

	return domain.Auth{
		AccessToken:  aws.ToString(authResult.AccessToken),
		IdToken:      aws.ToString(authResult.IdToken),
		RefreshToken: aws.ToString(authResult.RefreshToken),
	}, nil
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
	return domain.Auth{}, err
}

	authResult = output.AuthenticationResult

	return domain.Auth{
		AccessToken:  aws.ToString(authResult.AccessToken),
		IdToken:      aws.ToString(authResult.IdToken),
		RefreshToken: aws.ToString(authResult.RefreshToken),
	}, err
}