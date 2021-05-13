// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"
	"strconv"

	"github.com/globocom/gitlab-lint/db"
	"github.com/globocom/gitlab-lint/rules"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// projects godoc
// @Summary Show projects
// @Description get all projects
// @ID get-projects
// @Accept json
// @Produce json
// @Param q query string false "fuzzy search projects"
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

	project := &rules.Project{}

	searchStr := c.QueryParam("q")

	filter := db.FindFilter{
		Page: page,
		Sort: db.SortOption{
			Field: "pathwithnamespace",
			Order: db.SortAscending,
		},
		Query: db.BuildSearchQueryFromString(project, searchStr),
	}

	data, err := s.db.GetAll(project, filter)
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

	filter := db.FindFilter{
		Sort: db.SortOption{
			Field: "pathWithNamespace",
			Order: db.SortAscending,
		},
		Query: bson.M{"projectId": idInt},
	}

	projectRules, err := s.db.GetAll(&rules.Rule{}, filter)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"project": project,
		"rules":   projectRules,
	}

	return c.JSON(http.StatusOK, data)
}
