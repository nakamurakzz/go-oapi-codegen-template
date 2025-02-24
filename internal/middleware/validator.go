package middleware

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type RequestValidator struct {
	validator *validator.Validate
}

func NewRequestValidator() *RequestValidator {
	return &RequestValidator{validator: validator.New()}
}

func (rv *RequestValidator) Validate(i interface{}) error {
	if err := rv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, formatValidationError(err))
	}
	return nil
}

func formatValidationError(err error) map[string]interface{} {
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		errMessages := make(map[string]interface{})
		for _, e := range validationErrs {
			errMessages[e.Field()] = fmt.Sprintf(
				"Field validation for '%s' failed on the '%s' tag",
				e.Field(),
				e.Tag(),
			)
		}
		return errMessages
	}
	return map[string]interface{}{"message": err.Error()}
}
