// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"
	"sort"

	"github.com/globocom/gitlab-lint/db"
	"github.com/globocom/gitlab-lint/rules"
	"github.com/labstack/echo/v4"
)

// levels godoc
// @Summary Show levels
// @Description get levels
// @ID get-levels
// @Accept json
// @Produce json
// @Success 200 {object} Response{Data=map[string]int}
// @Router /levels [get]
func (s *server) levels(c echo.Context) error {
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

	levels := []rules.Level{}

	if len(statsAll) <= 0 {
		return c.JSON(http.StatusOK, levels)
	}

	stats := statsAll[0].(*rules.Stats)

	for levelID, levelCount := range stats.LevelsCount {
		level := rules.AllSeverities[levelID]
		level.Total = int32(levelCount)
		levels = append(levels, level)
	}

	// FIXME sort via database?
	sort.Slice(levels, func(i, j int) bool {
		return levels[i].Name < levels[j].Name
	})

	response := Response{
		Meta: MetaResponse{
			CurrentPage:  1,
			PerPage:      1,
			TotalOfItems: 1,
			TotalOfPages: 1,
		},
		Data: levels,
	}

	return c.JSON(http.StatusOK, response)
}
