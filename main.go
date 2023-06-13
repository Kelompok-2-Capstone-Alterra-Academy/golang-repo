package main

import (
	"capston-lms/internal/adapters/http"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	e := http.InitRoutes()
	e.Debug = true

	// Middleware CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Ganti dengan origin yang diizinkan
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	address := fmt.Sprintf(":%s", viper.GetString("EXSPOSE_PORT"))
	if address == ":" {
		address = ":8080" // Port default 8080 jika PORT tidak diset
	}
	// Start server
	e.Logger.Fatal(e.Start(address))
}
