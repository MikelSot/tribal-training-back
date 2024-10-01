package model

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Api              *fiber.App
	ProxyRouteSearch string
	ProxyRouteAuth   string
}
