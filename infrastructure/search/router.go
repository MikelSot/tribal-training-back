package search

import (
	"github.com/gofiber/fiber/v2"

	tribalAuth "github.com/MikelSot/tribal-training-auth/bootstrap"
	"github.com/MikelSot/tribal-training-back/model"
)

const (
	_privateRoutePrefix = "search"
)

func NewRouter(spec model.Config) {
	handler := buildHandler(spec)

	privateRoutes(spec.Api, handler, tribalAuth.ValidateJWT)
}

func buildHandler(spec model.Config) handler {
	return newHandler(spec.ProxyRouteSearch)
}

func privateRoutes(app *fiber.App, handler handler, middlewares ...fiber.Handler) {
	api := app.Group(_privateRoutePrefix, middlewares...)

	api.All("", handler.Proxy)
}
