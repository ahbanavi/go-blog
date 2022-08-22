package exceptions

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func JSONErrorHandler(c *fiber.Ctx, err error) error {
	// Default 500 statuscode
	code := fiber.StatusInternalServerError

	if e, ok := err.(validator.ValidationErrors); ok {
		var errors []*ErrorResponse
		for _, err := range e {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Status(code).JSON(err)
}
