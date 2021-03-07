// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"
	"strconv"

	"github.com/globocom/gitlab-lint/rules"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// projects godoc
// @Summary Show projects
// @Description get all projects
// @ID get-projects
// @Accept json
// @Produce json
// @Success 200 {array} rules.Project
// @Router /projects [get]
func (s *server) projects(c echo.Context) error {
	pageStr := ""
	pageStr = c.QueryParam("page")
	if pageStr == "" {
		pageStr = "1"
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return err
	}

	optProjects := options.Find()
	optProjects.SetSort(
		bson.D{primitive.E{Key: "pathwithnamespace", Value: 1}},
	)
	perPage := viper.GetInt("db.perPage")
	optProjects.SetSkip(int64((page - 1) * perPage))
	optProjects.SetLimit(int64(perPage))

	data, err := s.db.GetAll(&rules.Project{}, nil, optProjects)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, data)
}

// projects godoc
// @Summary Show project by id
// @Description get project by ID
// @ID get-projects-by-id
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} map[string]interface{}
// @Router /projects/{id} [get]
func (s *server) projectById(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	query := bson.M{"id": idInt}
	project := &rules.Project{}
	if err := s.db.Get(project, query, nil); err != nil {
		return err
	}

	optRules := options.Find().SetSort(
		bson.D{primitive.E{Key: "pathWithNamespace", Value: 1}},
	)
	queryRules := bson.M{"projectId": idInt}
	projectRules, err := s.db.GetAll(&rules.Rule{}, queryRules, optRules)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"project": project,
		"rules":   projectRules,
	}

	return c.JSON(http.StatusOK, data)
}
