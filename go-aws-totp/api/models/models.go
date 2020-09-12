package models

import "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

type SignUpJsonInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ConfirmSignUpJsonInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type ConfirmLoginJson struct {
	AuthInfo *cognitoidentityprovider.AdminInitiateAuthOutput `json:"auth_info"`
	Username string                                           `json:"username"`
	Token    string                                           `json:"token"`
}

type TotpJson struct {
	AccessToken string `json:"accessToken"`
	Code        string `json:"code"`
}
