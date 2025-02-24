package handler

import (
	"github.com/nakamurakzz/go-oapi-codegen-template/internal/gen/api"
	"github.com/nakamurakzz/go-oapi-codegen-template/internal/middleware"
)

type Server struct {
	validator *middleware.RequestValidator
}

var _ api.ServerInterface = &Server{}

func NewServer() Server {
	return Server{
		validator: middleware.NewRequestValidator(),
	}
}
