// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause Licens

package api

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func (s *server) healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "WORKING")
}
