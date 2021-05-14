// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"math"

	"github.com/globocom/gitlab-lint/db"
	"github.com/globocom/gitlab-lint/rules"
)

type MetaResponse struct {
	CurrentPage  int `json:"currentPage"`
	PerPage      int `json:"perPage"`
	TotalOfItems int `json:"totalOfItems"`
	TotalOfPages int `json:"totalOfPages"`
}

type Response struct {
	Data interface{}  `json:"data"`
	Meta MetaResponse `json:"meta"`
}

func (s *server) NewDataResponse(q rules.Queryable, filterOptions db.FindFilter, data interface{}) (Response, error) {
	response := Response{
		Meta: MetaResponse{
			CurrentPage: filterOptions.Page,
			PerPage:     filterOptions.PerPage,
		},
		Data: data,
	}

	totalOfItems, err := s.db.Count(q, filterOptions)
	if err != nil {
		return Response{}, err
	}

	response.Meta.TotalOfItems = totalOfItems

	div := float64(totalOfItems) / float64(filterOptions.PerPage)
	response.Meta.TotalOfPages = int(math.Ceil(div))

	return response, nil
}
