package main

import (
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"

	"github.com/marz32one/go-build-template/internal/config"
	"github.com/marz32one/go-build-template/internal/logging"
	api "github.com/marz32one/go-build-template/pkg/api/rest"
)

func main() {
	e := echo.New()

	// logger := zerolog.New(os.Stdout)
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
		LogStatus:   true,
		LogURI:      true,
		LogMethod:   true,
		LogRemoteIP: true,
		LogLatency:  true,
		LogError:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("method", v.Method).
				Str("URI", v.URI).
				Str("remote_ip", v.RemoteIP).
				Int("status", v.Status).
				Dur("latency", v.Latency).
				Err(v.Error).
				Msg("request")
			return nil
		},
	}))

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		logger.Info().
			Str("request", string(reqBody)).
			Msg("request body")
	}))

	// Initialize database
	config.Load()

	// Routes
	api.InitRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":9527"))
}
