package middleware

import (
	"strings"
	"warrantyapi/configuration"
	"warrantyapi/constant"
	"warrantyapi/model"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/golang-jwt/jwt/v5"

	"github.com/rs/zerolog/log"

	"github.com/gofiber/fiber/v2"
)

func GetUserNameFromToken(ctx *fiber.Ctx) string {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)
	return userName
}

func AuthenticateJWT(role string) func(*fiber.Ctx) error {
	// jwtSecret := config.Get("JWT_SECRET_KEY")
	return jwtware.New(jwtware.Config{
		// SigningKey: []byte(configuration.Secret),
		// SigningKey: SigningKey{Key: []byte(configuration.Secret)},
		SigningKey: jwtware.SigningKey{Key: []byte(configuration.Secret)},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			token := strings.TrimPrefix(ctx.Get("Authorization"), "Bearer ")
			log.Debug().Str("token", token).Send()
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			roles := claims["roles"].([]interface{})
			userName := claims["username"].(string)

			log.Debug().Str("userName", userName).Send()
			log.Debug().Any("role user", roles).Send()
			for _, roleInterface := range roles {
				roleMap := roleInterface.(map[string]interface{})
				if roleMap["role"] == role {
					return ctx.Next()
				}
			}

			return ctx.
				Status(fiber.StatusUnauthorized).
				JSON(model.GeneralResponse{
					Code:    constant.STATUS_CODE_UN_AUTHORIZED,
					Message: constant.MESSAGE_UN_AUTHORIZED,
					Data:    "",
				})
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.
					Status(fiber.StatusBadRequest).
					JSON(model.GeneralResponse{
						Code:    constant.STATUS_CODE_BAD_REQUEST,
						Message: constant.MESSAGE_BAD_REQUEST,
						// Data:    "",
						Data: "Missing or malformed JWT",
					})
			} else {
				return c.
					Status(fiber.StatusUnauthorized).
					JSON(model.GeneralResponse{
						Code:    constant.STATUS_CODE_UN_AUTHORIZED,
						Message: constant.MESSAGE_UN_AUTHORIZED,
						// Data:    "",
						Data: "Invalid or expired JWT",
					})
			}
		},
	})
}
