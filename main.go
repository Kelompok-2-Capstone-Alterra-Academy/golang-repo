package main

import (
	"capston-lms/internal/adapters/http"
	"fmt"
	"log"

	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	e := http.InitRoutes()
	e.Debug = true
	e.Pre(middleware.HTTPSRedirect())

	address := fmt.Sprintf(":%s", viper.GetString("EXSPOSE_PORT"))
	if address == ":" {
		address = ":8080" // Port default 8080 jika PORT tidak diset
	}
	// Load certificate and key from file
	certFile := "cert.pem"
	keyFile := "key.pem"
	err := e.StartTLS(address, certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}
}
