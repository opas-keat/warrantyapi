package common

import (
	"time"
	"warrantyapi/configuration"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string, roles []map[string]interface{}) string {
	// jwtSecret := config.Get("JWT_SECRET_KEY")
	// jwtExpired, err := strconv.Atoi(config.Get("JWT_EXPIRE_MINUTES_COUNT"))
	// exception.PanicLogging(err)

	claims := jwt.MapClaims{
		"username": username,
		"roles":    roles,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(configuration.Secret))
	if err != nil {
		println("GenerateToken %s\n", err)
	}

	return tokenSigned
}
