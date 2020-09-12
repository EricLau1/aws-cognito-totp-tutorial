package api

import (
	"flag"
	"go-aws-totp/admin"
	"go-aws-totp/api/controllers"
	"go-aws-totp/api/routes"
	"go-aws-totp/config"
	"go-aws-totp/provider"
	"go-aws-totp/totp"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

var port = flag.String("p", "8080", "api port")

func Run() {
	flag.Parse()
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on http://localhost:%s\n", *port)
	cfg := config.NewAwsConfig()
	prov := provider.NewAwsProvider(cfg)
	adm := admin.NewAwsAdmin(prov, cfg)
	awsTotp := totp.NewAwsTotp(prov, cfg)
	totpController := controllers.NewTotpController(awsTotp, adm)
	r := routes.NewRoutes(totpController)
	log.Fatal(http.ListenAndServe(":"+*port, r))
}
