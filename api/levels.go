// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"
	"sort"

	"github.com/globocom/gitlab-lint/rules"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// levels godoc
// @Summary Show levels
// @Description get levels
// @ID get-levels
// @Accept json
// @Produce json
// @Success 200 {array} rules.Level
// @Router /levels [get]
func (s *server) levels(c echo.Context) error {
	opts := options.Find()
	opts.SetSort(bson.D{primitive.E{Key: "_id", Value: -1}})
	opts.SetLimit(1)

	statsAll, err := s.db.GetAll(&rules.Stats{}, nil, opts)
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

	return c.JSON(http.StatusOK, levels)
}
