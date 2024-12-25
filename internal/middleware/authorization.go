package middleware

import (
	"author-service/pkg/exception"
	"errors"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/spf13/viper"
)

func Authorization() fiber.Handler {
	config := jwtware.Config{
		SigningKey:   []byte(viper.GetString("JWT_SECRET")),
		ErrorHandler: authError,
	}
	return jwtware.New(config)
}

func authError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return &exception.ErrWithCode{
			Code: fiber.StatusForbidden,
			Err:  errors.New("token.invalid"),
		}
	}

	// Return status 401 and failed authentication error.
	return &exception.ErrWithCode{
		Code: fiber.StatusForbidden,
		Err:  errors.New("token.invalid"),
	}
}
