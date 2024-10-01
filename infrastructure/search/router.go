package search

import (
	"github.com/gofiber/fiber/v2"

	"github.com/MikelSot/tribal-training-back/model"
)

const (
	_privateRoutePrefix = "search"
)

func NewRouter(spec model.Config) {
	handler := buildHandler(spec)

	// add middleware auth
	privateRoutes(spec.Api, handler)
}

func buildHandler(spec model.Config) handler {
	return newHandler(spec.ProxyRouteSearch)
}

func privateRoutes(app *fiber.App, handler handler, middlewares ...fiber.Handler) {
	api := app.Group(_privateRoutePrefix, middlewares...)

	api.All("", handler.Proxy)
}
