package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func GetLogValues(c echo.Context, v middleware.RequestLoggerValues) error {
	log.Info().
		Str("method", v.Method).
		Str("URI", v.URI).
		Str("remote_ip", v.RemoteIP).
		Int("status", v.Status).
		Dur("latency", v.Latency).
		Msg("request")
	return nil
}
