package main

import (
	_ "embed"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"math/rand"
	"strconv"
)

// jwtSecret is a secret key used to sign JWT tokens, which is defined in the config/config.json file.
// This key must be served securely, however, for the purposes of this tutorial, it is hardcoded here.
var jwtSecret = []byte("da01278c-7ae9-4570-be96-7bc7ebf4a441")

func main() {
	echoServer := echo.New()

	echoServer.Use(middleware.CORS())
	echoServer.Use(middleware.Logger())

	echoServer.GET("/auth/jwt", func(c echo.Context) error {
		userId := strconv.Itoa(rand.Intn(1000) + 1000)

		token, err := generateJwt(userId)

		if err != nil {
			echoServer.Logger.Error(err)

			return c.String(500, "Failed to generate JWT token")
		}

		return c.String(200, token)
	})

	echoServer.Logger.Error(echoServer.Start(":8001"))
}

// JWT token generation
func generateJwt(userId string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      userId,
		"channels": []string{"personal:broadcast"},
	}).SignedString(jwtSecret)
}
