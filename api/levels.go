// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
	"github.com/globocom/gitlab-lint/rules"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *server) levels(c echo.Context) error {
	pipeline := []bson.M{{"$group": bson.M{
		"_id":   "$level",
		"count": bson.M{"$sum": 1},
	}}}

	data, err := s.db.Aggregate(&rules.Rule{}, pipeline)
	if err != nil {
		return err
	}

	levels := []rules.Level{}
	for _, rule := range data {
		ruleId := rule["_id"].(string)
		level := rules.AllSeverities[ruleId]
		level.Total = rule["count"].(int32)
		levels = append(levels, level)
	}

	// FIXME sort via database?
	sort.Slice(levels, func(i, j int) bool {
		return levels[i].Name < levels[j].Name
	})

	return c.JSON(http.StatusOK, levels)
}
