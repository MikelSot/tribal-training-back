package search

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

type handler struct {
	proxyRouterSearch string
}

func newHandler(proxyRouterSearch string) handler {
	return handler{proxyRouterSearch}
}

func (h handler) Proxy(c *fiber.Ctx) error {
	url := fmt.Sprintf("%s%s", h.proxyRouterSearch, c.OriginalURL())
	if err := proxy.Do(c, url); err != nil {
		log.Warn("¡Uy! Fallo al enviar la solicitud a la API Search", err.Error())

		return c.Status(fiber.StatusBadGateway).JSON(`{"message": "¡Uy! Fallo al enviar la solicitud a la API Search"}`)
	}
	c.Response().Header.Del(fiber.HeaderServer)

	return nil
}
