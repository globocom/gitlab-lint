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

// rules godoc
// @Summary Show rules
// @Description get all projects
// @ID get-rules
// @Accept json
// @Produce json
// @Success 200 {array} interface{}
// @Router /rules [get]
func (s *server) rules(c echo.Context) error {
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

// rules godoc
// @Summary Show rule by id
// @Description get rule by ID
// @ID get-rules-by-id
// @Accept json
// @Produce json
// @Param id path string false "Rule ID"
// @Success 200 {object} map[string]interface{}
// @Router /rules/{id} [get]
func (s *server) rulesById(c echo.Context) error {
	id := c.Param("id")

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
