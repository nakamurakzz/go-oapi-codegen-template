package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nakamurakzz/go-oapi-codegen-template/internal/gen/api"
)

func (s Server) PostBooks(ctx echo.Context) error {
	var body api.PostBookRequest
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	err := s.validator.Validate(body)
	if err != nil {
		return err
	}

	// TODO: Usecase will be implemented here
	var description string
	if body.Description != nil {
		description = *body.Description
	}
	res := api.Book{
		Title:       body.Title,
		Author:      body.Author,
		Description: description,
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (s Server) GetBooks(ctx echo.Context) error {
	// Usecase will be implemented here
	books := []api.Book{
		{
			Title:       "Book 1",
			Author:      "Author 1",
			Description: "Description 1",
		},
		{
			Title:       "Book 2",
			Author:      "Author 2",
			Description: "Description 2",
		},
	}

	return ctx.JSON(http.StatusOK, books)
}
