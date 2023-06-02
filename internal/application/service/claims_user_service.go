package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetUserIDFromToken(e echo.Context) (int, error) {
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwt.MapClaims)
	id := int((*claims)["id"].(float64))
	return id, nil
}
