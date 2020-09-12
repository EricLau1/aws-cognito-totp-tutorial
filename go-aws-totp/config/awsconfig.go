package config

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"os"
)

const (
	AWS_USER_POOL_ID_ENV_KEY            = "AWS_USER_POOL_ID"
	AWS_USER_POOL_REGION_ENV_KEY        = "AWS_USER_POOL_REGION"
	AWS_USER_POOL_CLIENT_ID_ENV_KEY     = "AWS_USER_POOL_CLIENT_ID"
	AWS_ACCESS_KEY_ENV_KEY              = "AWS_ACCESS_KEY"
	AWS_SECRET_KEY_ENV_KEY              = "AWS_SECRET_KEY"
	AWS_USER_POOL_CLIENT_SECRET_ENV_KEY = "AWS_USER_POOL_CLIENT_SECRET"
)

type AwsConfig interface {
	UserPoolID() string
	UserPoolRegion() string
	UserPoolClientID() string
	AccessKey() string
	SecretKey() string
	UserPoolClientSecret() string
	SecretHash(string) string
}

type awsConfig struct {
}

func NewAwsConfig() AwsConfig {
	return &awsConfig{}
}

func (c *awsConfig) UserPoolID() string {
	return os.Getenv(AWS_USER_POOL_ID_ENV_KEY)
}

func (c *awsConfig) UserPoolRegion() string {
	return os.Getenv(AWS_USER_POOL_REGION_ENV_KEY)
}

func (c *awsConfig) UserPoolClientID() string {
	return os.Getenv(AWS_USER_POOL_CLIENT_ID_ENV_KEY)
}

func (c *awsConfig) UserPoolClientSecret() string {
	return os.Getenv(AWS_USER_POOL_CLIENT_SECRET_ENV_KEY)
}

func (c *awsConfig) AccessKey() string {
	return os.Getenv(AWS_ACCESS_KEY_ENV_KEY)
}

func (c *awsConfig) SecretKey() string {
	return os.Getenv(AWS_SECRET_KEY_ENV_KEY)
}

func (c *awsConfig) SecretHash(username string) string {
	hasher := hmac.New(sha256.New, []byte(c.UserPoolClientSecret()))

	hasher.Write([]byte(username))
	hasher.Write([]byte(c.UserPoolClientID()))

	hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	return hash
}
