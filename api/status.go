// Copyright (c) 2021
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"
	"time"

	"github.com/globocom/gitlab-lint/db"
	"github.com/labstack/echo/v4"
)

const (
	StatusWorking = "WORKING"
	StatusFailed  = "FAILED"
)

type Status struct {
	Project  string    `json:"project"`
	Version  string    `json:"version"`
	Status   string    `json:"status"`
	Services []Service `json:"services"`
}

type Service struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Elapsed string `json:"elapsed"`
	Message string `json:"message,omitempty"`
}

// status godoc
// @Summary Show status services
// @Description get status
// @ID get-status
// @Accept json
// @Produce json
// @Success 200 {object} Status
// @Router /status [get]
func (s *server) status(c echo.Context) error {
	serviceMongo := healthyMongo(s.db)
	services := []Service{
		serviceMongo,
	}

	status := StatusWorking
	for _, service := range services {
		if service.Status != StatusWorking {
			status = StatusFailed
			break
		}
	}

	data := Status{
		Project:  "gitlab-lint",
		Version:  Version,
		Status:   status,
		Services: services,
	}

	return c.JSON(http.StatusOK, data)
}

func healthyMongo(db db.DB) Service {
	service := Service{
		Name:   "mongo",
		Status: StatusWorking,
	}

	t1 := time.Now()
	if err := db.Ping(); err != nil {
		service.Status = StatusFailed
		service.Message = err.Error()
	}
	service.Elapsed = time.Since(t1).String()

	return service
}
