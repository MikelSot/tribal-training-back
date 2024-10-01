package bootstrap

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"

	"github.com/MikelSot/tribal-training-back/infrastructure/handler"
	"github.com/MikelSot/tribal-training-back/model"
)

func Run() {
	_ = godotenv.Load()

	app := newFiber()

	handler.InitRoutes(model.Config{
		Api:              app,
		ProxyRouteSearch: getProxyRouteSearch(),
		ProxyRouteAuth:   getProxyRouteAuth(),
	})

	log.Fatal(app.Listen(getPort()))
}
