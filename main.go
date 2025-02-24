package main

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nakamurakzz/go-oapi-codegen-template/handler"
	"github.com/nakamurakzz/go-oapi-codegen-template/internal/gen/api"
	mw "github.com/nakamurakzz/go-oapi-codegen-template/internal/middleware"
)

func main() {
	server := handler.NewServer()
	e := echo.New()
	e.Validator = mw.NewRequestValidator()
	e.HTTPErrorHandler = mw.CustomErrorHandler

	api.RegisterHandlers(e, server)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Start server
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}
