// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
	"github.com/globocom/gitlab-lint/rules"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *server) rules(c echo.Context) error {
	id := c.Param("id")

	if id != "" {
		query := bson.M{"ruleId": id}
		optRules := options.Find().SetSort(
			bson.D{primitive.E{Key: "pathWithNamespace", Value: 1}},
		)
		projects, err := s.db.GetAll(&rules.Rule{}, query, optRules)
		if err != nil {
			return err
		}
		ruler := rules.MyRegistry.RulesFn[id]
		data := map[string]interface{}{
			"rule":     ruler,
			"projects": projects,
		}
		return c.JSON(http.StatusOK, data)
	}

	data := []rules.Ruler{}
	for _, rule := range rules.MyRegistry.RulesFn {
		data = append(data, rule)
	}

	// FIXME sort via database?
	sort.Slice(data, func(i, j int) bool {
		return data[i].GetSlug() < data[j].GetSlug()
	})

	return c.JSON(http.StatusOK, data)
}
