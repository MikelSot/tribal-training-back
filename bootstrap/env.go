package bootstrap

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
)

const _nameAppDefault = "tribal-training-gateway"

const _portDefault = ":8080"

const (
	_allowOriginsDefault = "*"
	_allowMethodsDefault = "GET,POST"
)

const (
	_proxyRouteSearchDefault = "/search"
	_proxyRouteAuthDefault   = "/auth"
)

func getApplicationName() string {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		return _nameAppDefault
	}

	return appName
}

func getPort() string {
	port := os.Getenv("FIBER_PORT")
	if port == "" {
		return _portDefault
	}

	return port
}

func getAllowOrigins() string {
	allowedOrigins := os.Getenv("ALLOW_ORIGINS")
	if allowedOrigins == "" {
		return _allowOriginsDefault
	}

	return allowedOrigins
}

func getAllowMethods() string {
	allowedMethods := os.Getenv("ALLOW_METHODS")
	if allowedMethods == "" {
		return _allowMethodsDefault
	}

	return allowedMethods
}

func getProxyRouteSearch() string {
	proxyRouteQR := os.Getenv("PROXY_ROUTE_QR")
	if proxyRouteQR == "" {
		log.Warn("proxy route qr not found")

		return _proxyRouteSearchDefault
	}

	return proxyRouteQR
}

func getProxyRouteAuth() string {
	proxyRouteAuth := os.Getenv("PROXY_ROUTE_AUTH")
	if proxyRouteAuth == "" {
		log.Warn("proxy route auth not found")

		return _proxyRouteAuthDefault
	}

	return proxyRouteAuth
}

func getSignKey() string {
	signKey := os.Getenv("SIGN_KEY")
	if signKey == "" {
		log.Warn("sign key not found")

		return ""
	}

	return signKey
}
