package main

import (
	"capston-lms/internal/adapters/http"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func main() {
	e := http.InitRoutes()
	e.Debug = true

	secretJWT := viper.GetString("EXSPOSE_PORT")
	if secretJWT == "" {
		log.Fatal("EXSPOSE_PORT is not set")
	}

	address := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if address == ":" {
		address = ":8080" // Port default 8080 jika PORT tidak diset
	}

	err := e.Start(address)
	if err != nil {
		log.Fatal(err)
	}
}
