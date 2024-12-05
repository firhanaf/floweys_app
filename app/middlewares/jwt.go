package middlewares

import (
	"errors"
	"floweys_app/app/config"
	"github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"time"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(config.JWT_SECRET),
		SigningMethod: "HS256",
	})
}

func CreateToken(ID uint, Username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = ID
	claims["username"] = Username
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWT_SECRET))
}

func ExtractTokenUser(e echo.Context) (uint, string, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(float64)
		username := claims["username"].(string)
		return uint(userId), username, nil
	}
	return 0, "not found", errors.New("token invalid")
}
