// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/globocom/gitlab-lint/rules"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	opts := options.Find()
	opts.SetSort(bson.D{primitive.E{Key: "_id", Value: -1}})
	opts.SetLimit(1)

	statsAll, err := s.db.GetAll(&rules.Stats{}, nil, opts)
	if err != nil {
		return err
	}

	if len(statsAll) <= 0 {
		return c.JSON(http.StatusOK, rules.Stats{})
	}

	return c.JSON(http.StatusOK, statsAll[0])
}
