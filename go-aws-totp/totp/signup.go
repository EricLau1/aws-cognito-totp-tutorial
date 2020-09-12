package totp

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func (t *awsTotp) SignUp(username, email, password string) (*cognitoidentityprovider.SignUpOutput, error) {

	input := &cognitoidentityprovider.SignUpInput{
		SecretHash: aws.String(t.c.SecretHash(username)),
		ClientId:   aws.String(t.c.UserPoolClientID()),
		Password:   aws.String(password),
		Username:   aws.String(username),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			&cognitoidentityprovider.AttributeType{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
		},
	}

	return t.p.SignUp(input)
}

func (t *awsTotp) ConfirmSignUp(username, email, code string) (*cognitoidentityprovider.ConfirmSignUpOutput, error) {

	input := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(t.c.UserPoolClientID()),
		ConfirmationCode: aws.String(code),
		SecretHash:       aws.String(t.c.SecretHash(username)),
		Username:         aws.String(username),
	}

	output, err := t.p.ConfirmSignUp(input)
	if err != nil {
		return nil, err
	}

	_, err = t.SetPreferredUsername(username, email)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (t *awsTotp) SetPreferredUsername(currentUsername, preferredUsername string) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error) {

	input := &cognitoidentityprovider.AdminUpdateUserAttributesInput{
		UserPoolId: aws.String(t.c.UserPoolID()),
		Username:   aws.String(currentUsername),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			&cognitoidentityprovider.AttributeType{
				Name:  aws.String(cognitoidentityprovider.AliasAttributeTypePreferredUsername),
				Value: aws.String(preferredUsername),
			},
		},
	}

	return t.p.AdminUpdateUserAttributes(input)
}
