package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marz32one/go-build-template/pkg/api"
	"github.com/rs/zerolog/log"
)

func ErrorHandlingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			// Extract HTTPError, if possible
			code := http.StatusInternalServerError
			msg := err.Error()
			if he, ok := err.(*echo.HTTPError); ok {
				code = he.Code
				msg = he.Message.(string)
			}

			// Log the error using zerolog
			log.Error().
				Err(err).
				Str("method", c.Request().Method).
				Str("path", c.Path()).
				Str("remote_ip", c.RealIP()).
				Msg("API error")

			// Send the custom error response
			return api.SendErrorResponse(c, code, msg)
		}
		return nil
	}
}
