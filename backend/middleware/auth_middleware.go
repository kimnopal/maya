package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kimnopal/maya/converter"
	"github.com/kimnopal/maya/model"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	header := ctx.Get("Authorization")
	if header == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(converter.ToWebResponse(
			fiber.StatusUnauthorized,
			fiber.ErrUnauthorized.Message,
			nil,
		))
	}

	if !strings.Contains(header, "Bearer") {
		return ctx.Status(fiber.StatusBadRequest).JSON(converter.ToWebResponse(
			fiber.StatusBadRequest,
			fiber.ErrBadRequest.Message,
			nil,
		))
	}

	token := strings.Replace(header, "Bearer ", "", -1)

	userClaims := new(model.UserClaims)
	jwtToken, err := jwt.ParseWithClaims(token, userClaims, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrBadRequest
		} else if method != jwt.SigningMethodHS256 {
			return nil, fiber.ErrBadRequest
		}

		return []byte("rahasia"), nil
	})

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(converter.ToWebResponse(
			fiber.StatusUnauthorized,
			fiber.ErrUnauthorized.Message,
			nil,
		))
	}

	if !jwtToken.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(converter.ToWebResponse(
			fiber.StatusUnauthorized,
			fiber.ErrUnauthorized.Message,
			nil,
		))
	}

	ctx.Locals("auth", userClaims.UserResponse)

	return ctx.Next()
}
