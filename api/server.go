// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/globocom/gitlab-lint/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const Version = "0.1.0"

type Server interface {
	http.Handler
	Start()
}

type server struct {
	*echo.Echo
	db db.DB
}

func (s *server) Start() {
	addr := fmt.Sprintf(":%d", viper.GetInt("port"))
	s.Echo.Logger.Fatal(s.Echo.Start(addr))
}

func NewServer() (Server, error) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: viper.GetString("sentry.dsn"),
	}); err != nil {
		log.Warnf("Sentry initialization failed: %v", err)
	}

	echoInstance := echo.New()
	echoInstance.HideBanner = true
	echoInstance.HTTPErrorHandler = httpErrorHandler

	dbInstance, err := db.NewMongoSession()
	if err != nil {
		return nil, err
	}

	server := &server{
		Echo: echoInstance,
		db:   dbInstance,
	}

	echoInstance.Pre(middleware.RemoveTrailingSlash())

	echoInstance.Use(
		middleware.LoggerWithConfig(
			middleware.LoggerConfig{
				Format: "timestamp=${time_rfc3339} " +
					"method=${method} " +
					"request_uri=${uri} " +
					"status=${status} " +
					"request_id=${id} " +
					"latency=${latency_human}\n",
			},
		),
	)
	echoInstance.Use(middleware.Recover())
	echoInstance.Use(sentryecho.New(sentryecho.Options{
		Repanic:         true,
		WaitForDelivery: false,
		Timeout:         viper.GetDuration("sentry.timeout") * time.Second,
	}))

	echoInstance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			echo.GET, echo.OPTIONS,
		},
	}))

	echoInstance.GET("/", server.index)
	echoInstance.GET("/api/v1/projects", server.projects)
	echoInstance.GET("/api/v1/projects/:id", server.projectById)
	echoInstance.GET("/api/v1/rules", server.rules)
	echoInstance.GET("/api/v1/rules/:id", server.rulesById)
	echoInstance.GET("/api/v1/levels", server.levels)
	echoInstance.GET("/api/v1/stats", server.stats)
	echoInstance.GET("/healthcheck", server.healthcheck)
	echoInstance.GET("/swagger/*", echoSwagger.WrapHandler)

	return server, nil
}
