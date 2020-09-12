package controllers

import (
	"encoding/json"
	"go-aws-totp/admin"
	"go-aws-totp/api/models"
	"go-aws-totp/api/server"
	"go-aws-totp/api/utils"
	"go-aws-totp/totp"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type TotpController interface {
	PostSignUp(w http.ResponseWriter, r *http.Request)
	PostConfirmSignUp(w http.ResponseWriter, r *http.Request)
	PostLogin(w http.ResponseWriter, r *http.Request)
	PostConfirmLogin(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	PostAssociateToken(w http.ResponseWriter, r *http.Request)
	PostConfirmToken(w http.ResponseWriter, r *http.Request)
}

type totpController struct {
	cli totp.AwsTotp
	adm admin.AwsAdmin
}

func NewTotpController(cli totp.AwsTotp, adm admin.AwsAdmin) TotpController {
	return &totpController{cli: cli, adm: adm}
}

func (c *totpController) PostSignUp(w http.ResponseWriter, r *http.Request) {
	input := new(models.SignUpJsonInput)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		server.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, input)
	if err != nil {
		server.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}
	input.Email = utils.NormalizeEmail(input.Email)
	output, err := c.cli.SignUp(input.Email, input.Email, input.Password)
	if err != nil {
		server.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}
	server.RespondWithJson(w, http.StatusCreated, output)
}

func (c *totpController) PostConfirmSignUp(w http.ResponseWriter, r *http.Request) {
	input := new(models.ConfirmSignUpJsonInput)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		server.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, input)
	if err != nil {
		server.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}
	input.Email = utils.NormalizeEmail(input.Email)
	user, err := c.adm.GetByEmail(input.Email)
	if err != nil {
		server.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	input.Username = *user.Username
	output, err := c.cli.ConfirmSignUp(input.Username, input.Email, input.Code)
	if err != nil {
		server.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}
	server.RespondWithJson(w, http.StatusOK, output)
}

func (c *totpController) PostLogin(w http.ResponseWriter, r *http.Request) {
	input := new(models.SignUpJsonInput)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		server.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, input)
	if err != nil {
		server.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}
	input.Email = utils.NormalizeEmail(input.Email)
	output, err := c.cli.Login(input.Email, input.Password)
	if err != nil {
		server.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	server.RespondWithJson(w, http.StatusOK, output)
}

func (c *totpController) PostConfirmLogin(w http.ResponseWriter, r *http.Request) {
	input := new(models.ConfirmLoginJson)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		server.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, input)
	if err != nil {
		server.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}
	input.Username = utils.NormalizeEmail(input.Username)
	input.Token = strings.TrimSpace(input.Token)
	output, err := c.cli.ConfimLogin(input.AuthInfo, input.Username, input.Token)
	if err != nil {
		server.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	server.RespondWithJson(w, http.StatusOK, output)
}

func (c *totpController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	user, err := c.adm.GetUser(username)
	if err != nil {
		server.RespondWithError(w, http.StatusBadRequest, err)
	} else {
		server.RespondWithJson(w, http.StatusOK, user)
	}
}

func (c *totpController) PostAssociateToken(w http.ResponseWriter, r *http.Request) {
	input := new(models.TotpJson)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		server.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, input)
	if err != nil {
		server.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}
	output, err := c.cli.InitTotp(input.AccessToken)
	if err != nil {
		server.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	server.RespondWithJson(w, http.StatusOK, output)
}

func (c *totpController) PostConfirmToken(w http.ResponseWriter, r *http.Request) {
	input := new(models.TotpJson)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		server.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, input)
	if err != nil {
		server.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}
	output, err := c.cli.ConfirmTotp(input.AccessToken, input.Code)
	if err != nil {
		server.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	server.RespondWithJson(w, http.StatusOK, output)
}
