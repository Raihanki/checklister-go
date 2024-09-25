package middleware

import (
	"strings"

	"github.com/Raihanki/checklisters/internal/helpers"
	"github.com/gofiber/fiber/v2"
)

func Authenticate(ctx *fiber.Ctx) error {
	authorizationHeader := ctx.Get("Authorization")
	if authorizationHeader == "" {
		return helpers.JsonResponse(ctx, fiber.StatusUnauthorized, "Unauthorized", nil)
	}

	bearerToken := strings.Split(authorizationHeader, " ")[1]
	if bearerToken == "" {
		return helpers.JsonResponse(ctx, fiber.StatusUnauthorized, "Unauthorized", nil)
	}

	validatedToken, errValidatedToken := helpers.ValidateToken(bearerToken)
	if errValidatedToken != nil {
		return helpers.JsonResponse(ctx, fiber.StatusUnauthorized, "Unauthorized", nil)
	}

	userId, err := validatedToken.GetSubject()
	if err != nil {
		return helpers.JsonResponse(ctx, fiber.StatusUnauthorized, "Unauthorized", nil)
	}

	ctx.Locals("user_id", userId)

	return ctx.Next()
}
