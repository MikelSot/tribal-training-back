package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

type handler struct {
	proxyRouterAuth string
}

func newHandler(proxyRouterAuth string) handler {
	return handler{proxyRouterAuth}
}

func (h handler) Proxy(c *fiber.Ctx) error {
	url := fmt.Sprintf("%s%s", h.proxyRouterAuth, c.OriginalURL())
	log.Warn("alsndclsjnd")
	log.Warn(url)

	if err := proxy.Do(c, url); err != nil {
		log.Warn("¡Uy! Fallo al enviar la solicitud a la API Auth: ", err.Error())

		return c.Status(fiber.StatusBadGateway).JSON(`{"errors":[{"code":500,"message":"¡Uy! Fallo al enviar la solicitud a la API Auth"}]}`)
	}
	c.Response().Header.Del(fiber.HeaderServer)

	return nil
}
