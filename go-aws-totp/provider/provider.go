package provider

import (
	"go-aws-totp/config"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func NewAwsProvider(cfg config.AwsConfig) *cognitoidentityprovider.CognitoIdentityProvider {

	creds := credentials.NewStaticCredentials(cfg.AccessKey(), cfg.SecretKey(), "")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(cfg.UserPoolRegion()),
		Credentials: creds,
	})

	if err != nil {
		log.Fatal(err)
	}

	return cognitoidentityprovider.New(sess)
}
