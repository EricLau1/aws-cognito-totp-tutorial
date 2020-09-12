package totp

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type TotpConfig interface {
	UserPoolID() string
	UserPoolRegion() string
	UserPoolClientID() string
	SecretHash(string) string
}

type AwsTotp interface {
	SignUp(username string, email string, password string) (*cognitoidentityprovider.SignUpOutput, error)
	ConfirmSignUp(username, email, code string) (*cognitoidentityprovider.ConfirmSignUpOutput, error)
	SetPreferredUsername(currentUsername, preferredUsername string) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error)
	Login(email string, password string) (*cognitoidentityprovider.AdminInitiateAuthOutput, error)
	ConfimLogin(initAuthOutput *cognitoidentityprovider.AdminInitiateAuthOutput, username, token string) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error)
	InitTotp(accessToken string) (*cognitoidentityprovider.AssociateSoftwareTokenOutput, error)
	ConfirmTotp(accessToken, code string) (*cognitoidentityprovider.VerifySoftwareTokenOutput, error)
	EnableTotp(accessToken string) (*cognitoidentityprovider.SetUserMFAPreferenceOutput, error)
}

type awsTotp struct {
	p *cognitoidentityprovider.CognitoIdentityProvider
	c TotpConfig
}

func NewAwsTotp(p *cognitoidentityprovider.CognitoIdentityProvider, cfg TotpConfig) AwsTotp {
	return &awsTotp{p: p, c: cfg}
}

func (t *awsTotp) InitTotp(accessToken string) (*cognitoidentityprovider.AssociateSoftwareTokenOutput, error) {

	input := &cognitoidentityprovider.AssociateSoftwareTokenInput{
		AccessToken: aws.String(accessToken),
	}

	return t.p.AssociateSoftwareToken(input)
}

func (t *awsTotp) ConfirmTotp(accessToken, code string) (*cognitoidentityprovider.VerifySoftwareTokenOutput, error) {

	input := &cognitoidentityprovider.VerifySoftwareTokenInput{
		FriendlyDeviceName: aws.String("totp"),
		AccessToken:        aws.String(accessToken),
		UserCode:           aws.String(code),
	}

	output, err := t.p.VerifySoftwareToken(input)
	if err != nil {
		return nil, err
	}

	enabled, err := t.EnableTotp(accessToken)
	if err != nil {
		return nil, err
	} else {
		log.Println("EnableTotp: ", enabled)
	}

	return output, nil
}

func (t *awsTotp) EnableTotp(accessToken string) (*cognitoidentityprovider.SetUserMFAPreferenceOutput, error) {

	input := &cognitoidentityprovider.SetUserMFAPreferenceInput{
		AccessToken: aws.String(accessToken),
		SoftwareTokenMfaSettings: &cognitoidentityprovider.SoftwareTokenMfaSettingsType{
			Enabled:      aws.Bool(true),
			PreferredMfa: aws.Bool(true),
		},
	}

	return t.p.SetUserMFAPreference(input)
}
