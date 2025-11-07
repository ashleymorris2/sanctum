package middleware

import (
	"errors"
	"fmt"
	"metrics/internal/auth"
	"metrics/internal/dto"
	"metrics/internal/validators"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func JSONErrorHandler(defaultHandler echo.HTTPErrorHandler) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		resp := dto.ErrorResponse{
			Code:    "INTERNAL_ERROR",
			Message: "Unexpected error",
		}

		var vErrs validator.ValidationErrors
		if errors.As(err, &vErrs) {
			code = http.StatusBadRequest
			resp.Code = "VALIDATION_FAILED"
			resp.Message = "One or more fields are invalid"
			resp.Details = validators.FormatErrors(err)
		}

		var apiErr auth.APIError
		if errors.As(err, &apiErr) {
			code = apiErr.HTTPStatus()
			resp.Code = apiErr.Code()
			resp.Message = apiErr.Message()
		}

		var reqErr *echo.HTTPError
		if errors.As(err, &reqErr) {
			code = reqErr.Code
			switch msg := reqErr.Message.(type) {
			case string:
				resp.Message = msg
			case *dto.ErrorResponse:
				resp = *msg
			default:
				resp.Message = fmt.Sprint(msg)
			}
		}

		if !c.Response().Committed {
			_ = c.JSON(code, resp)
		}
	}
}
