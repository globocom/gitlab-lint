// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *server) index(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"name":    "gitlab-lint API",
		"version": Version,
	})
}
