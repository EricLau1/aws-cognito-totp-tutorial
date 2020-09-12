package admin

import (
	"fmt"
	"go-aws-totp/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type AwsAdmin interface {
	GetByEmail(email string) (*cognitoidentityprovider.UserType, error)
	GetUser(username string) (*cognitoidentityprovider.AdminGetUserOutput, error)
	DeleteByEmail(email string) (*cognitoidentityprovider.AdminDeleteUserOutput, error)
}

type awsAdmin struct {
	p *cognitoidentityprovider.CognitoIdentityProvider
	c config.AwsConfig
}

func NewAwsAdmin(p *cognitoidentityprovider.CognitoIdentityProvider, c config.AwsConfig) AwsAdmin {
	return &awsAdmin{p: p, c: c}
}

func (a *awsAdmin) GetUser(username string) (*cognitoidentityprovider.AdminGetUserOutput, error) {
	input := &cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: aws.String(a.c.UserPoolID()),
		Username:   aws.String(username),
	}

	return a.p.AdminGetUser(input)
}

func (a *awsAdmin) GetByEmail(email string) (*cognitoidentityprovider.UserType, error) {

	f := fmt.Sprintf("\"email\" = \"%s\"", email)

	input := &cognitoidentityprovider.ListUsersInput{
		Limit:      aws.Int64(1),
		Filter:     aws.String(f),
		UserPoolId: aws.String(a.c.UserPoolID()),
	}

	output, err := a.p.ListUsers(input)
	if err != nil {
		return nil, err
	}
	if len(output.Users) == 0 {
		return nil, fmt.Errorf("%s not found", email)
	}
	return output.Users[0], nil
}

func (a *awsAdmin) DeleteByEmail(email string) (*cognitoidentityprovider.AdminDeleteUserOutput, error) {

	user, err := a.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	input := &cognitoidentityprovider.AdminDeleteUserInput{
		UserPoolId: aws.String(a.c.UserPoolID()),
		Username:   user.Username,
	}

	return a.p.AdminDeleteUser(input)
}
