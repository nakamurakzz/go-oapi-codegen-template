package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nakamurakzz/go-oapi-codegen-template/internal/gen/api"
)

func CustomErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	switch e := err.(type) {
	case *echo.HTTPError:
		code = e.Code
		msg = e.Message
	default:
		msg = err.Error()
	}

	// エラーレスポンスの構造化
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, api.ErrorResponse{
				Status:  code,
				Message: fmt.Sprint(msg),
			})
		}
	}
}
