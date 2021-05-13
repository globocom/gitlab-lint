// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/globocom/gitlab-lint/db"
	"github.com/globocom/gitlab-lint/rules"
	"github.com/labstack/echo/v4"
)

// stats godoc
// @Summary Show stats
// @Description get stats
// @ID get-stats
// @Accept json
// @Produce json
// @Success 200 {object} rules.Stats
// @Router /stats [get]
func (s *server) stats(c echo.Context) error {

	filter := CreateFilterFromQueryParam(&rules.Stats{}, c.QueryParams())
	filter.PerPage = 1
	filter.Sort = db.SortOption{
		Field: "_id",
		Order: db.SortDescending,
	}

	statsAll, err := s.db.GetAll(&rules.Stats{}, filter)
	if err != nil {
		return err
	}

	if len(statsAll) <= 0 {
		return c.JSON(http.StatusOK, rules.Stats{})
	}

	return c.JSON(http.StatusOK, statsAll[0])
}
