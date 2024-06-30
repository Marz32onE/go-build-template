package main

import (
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"

	"github.com/marz32one/go-build-template/internal/config"
	api "github.com/marz32one/go-build-template/pkg/api/rest"
	customMiddleware "github.com/marz32one/go-build-template/pkg/middleware"
	"github.com/marz32one/go-build-template/pkg/util/logging"
)

func main() {
	// Initialize database
	config.Load()

	e := echo.New()

	logLevel, err := strconv.Atoi(os.Getenv("DEBUG_LEVEL"))
	if err == nil {
		logLevel = 0
	}
	logger := logging.Configure(logging.Config{
		LogLevel:              zerolog.Level(logLevel),
		ConsoleLoggingEnabled: true,
		FileLoggingEnabled:    true,
		Directory:             "./logs/",
		Filename:              "go-template.log",
		MaxSize:               10,
		MaxBackups:            1,
		MaxAge:                28,
	})

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:     true,
		LogURI:        true,
		LogMethod:     true,
		LogRemoteIP:   true,
		LogLatency:    true,
		LogError:      true,
		LogValuesFunc: customMiddleware.GetLogValues,
	}))

	// Middleware
	// e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		logger.Info().
			Str("request", string(reqBody)).
			Msg("request body")
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			logger.Error().Err(err).Bytes("stack", stack).Msg("Panic recovered")
			return err
		},
	}))

	// Timeout middleware with configuration
	timeoutConfig := middleware.TimeoutConfig{
		Timeout: 10 * time.Second,
	}
	e.Use(middleware.TimeoutWithConfig(timeoutConfig))

	e.Use(customMiddleware.AuditMiddlewareWithConfig())

	// Routes
	api.InitRoutes(e)
	e.GET("/error", func(c echo.Context) error {
		logger.Error().Msg("Error endpoint called")
		return echo.NewHTTPError(500, "Internal Server Error")
	})

	// Start server
	e.Logger.Fatal(e.Start(":9527"))
}
