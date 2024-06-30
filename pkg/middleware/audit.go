package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// AuditMiddlewareWithConfig returns a middleware that checks for the X-APIKEY header and logs the request details.
func AuditMiddlewareWithConfig() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			envAPIKey := os.Getenv("API-KEY")
			reqAPIKey := c.Request().Header.Get("X-APIKEY")
			if os.Getenv("ENV") == "PROD" && reqAPIKey != envAPIKey {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid API Key")
			}

			// Log the request details
			log.Printf("Method: %s, URI: %s, IP: %s, UserAgent: %s, Duration: %v\n",
				c.Request().Method,
				c.Request().RequestURI,
				c.Request().RemoteAddr,
				c.Request().UserAgent(),
				time.Since(start),
			)

			return next(c)
		}
	}
}
