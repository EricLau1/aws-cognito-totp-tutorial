package totp

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func (t *awsTotp) Login(email string, password string) (*cognitoidentityprovider.AdminInitiateAuthOutput, error) {

	params := map[string]*string{
		"USERNAME":    &email,
		"PASSWORD":    &password,
		"SECRET_HASH": aws.String(t.c.SecretHash(email)),
	}

	input := &cognitoidentityprovider.AdminInitiateAuthInput{
		AuthFlow:       aws.String("ADMIN_USER_PASSWORD_AUTH"),
		ClientId:       aws.String(t.c.UserPoolClientID()),
		UserPoolId:     aws.String(t.c.UserPoolID()),
		AuthParameters: params,
	}

	return t.p.AdminInitiateAuth(input)
}

// https://docs.amazonaws.cn/cognito-user-identity-pools/latest/APIReference/API_RespondToAuthChallenge.html
func (t *awsTotp) ConfimLogin(initAuthOutput *cognitoidentityprovider.AdminInitiateAuthOutput, username, token string) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error) {

	if *initAuthOutput.ChallengeName != "SOFTWARE_TOKEN_MFA" {
		return nil, errors.New("ChallengeName must be: SOFTWARE_TOKEN_MFA")
	}

	responses := map[string]*string{
		"USERNAME":                &username,
		"SOFTWARE_TOKEN_MFA_CODE": &token,
		"SECRET_HASH":             aws.String(t.c.SecretHash(username)),
	}

	input := &cognitoidentityprovider.AdminRespondToAuthChallengeInput{
		ChallengeName:      initAuthOutput.ChallengeName,
		Session:            initAuthOutput.Session,
		ChallengeResponses: responses,
		ClientId:           aws.String(t.c.UserPoolClientID()),
		UserPoolId:         aws.String(t.c.UserPoolID()),
	}

	return t.p.AdminRespondToAuthChallenge(input)
}
