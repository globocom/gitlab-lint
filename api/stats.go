// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/globocom/gitlab-lint/rules"
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
// @Success 200 {array} rules.Stats
// @Router /stats [get]
func (s *server) stats(c echo.Context) error {
	opts := options.FindOne().SetSort(
		bson.D{primitive.E{Key: "$natural", Value: -1}},
	)
	stats := &rules.Stats{}
	if err := s.db.Get(stats, nil, opts); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, stats)
}
