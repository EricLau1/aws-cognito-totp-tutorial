package routes

import (
	"go-aws-totp/api/controllers"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRoutes(totpController controllers.TotpController) http.Handler {
	r := mux.NewRouter().StrictSlash(true)

	r.Path("/api/totp/signup").Methods(http.MethodPost).HandlerFunc(totpController.PostSignUp)
	r.Path("/api/totp/confirm-signup").Methods(http.MethodPost).HandlerFunc(totpController.PostConfirmSignUp)
	r.Path("/api/totp/login").Methods(http.MethodPost).HandlerFunc(totpController.PostLogin)
	r.Path("/api/totp/confirm-login").Methods(http.MethodPost).HandlerFunc(totpController.PostConfirmLogin)
	r.Path("/api/totp/user/{username}").Methods(http.MethodGet).HandlerFunc(totpController.GetUser)
	r.Path("/api/totp/associate").Methods(http.MethodPost).HandlerFunc(totpController.PostAssociateToken)
	r.Path("/api/totp/confirm").Methods(http.MethodPost).HandlerFunc(totpController.PostConfirmToken)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "Accept", "X-Request"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	return handlers.CORS(headers, methods, origins)(r)
}
