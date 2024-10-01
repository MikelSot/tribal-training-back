package auth

import (
	"github.com/gofiber/fiber/v2"

	"github.com/MikelSot/tribal-training-back/model"
)

const (
	_publicRoutePrefix = "auth"
)

func NewRouter(spec model.Config) {
	handler := buildHandler(spec)

	publicRoutes(spec.Api, handler)
}

func buildHandler(spec model.Config) handler {
	return newHandler(spec.ProxyRouteAuth)
}

func publicRoutes(app *fiber.App, handler handler, middlewares ...fiber.Handler) {
	api := app.Group(_publicRoutePrefix, middlewares...)

	api.Post("/api/v1/login", handler.Proxy)
	api.Post("/api/v1/register", handler.Proxy)
}
