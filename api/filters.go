package api

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/globocom/gitlab-lint/db"
	"github.com/globocom/gitlab-lint/rules"
)

func parseIntFromString(str string) int {
	number, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return number
}

func getSortFromString(str string) int {
	lowerStr := strings.ToLower(str)
	sortMap := map[string]int{
		"asc":  db.SortAscending,
		"1":    db.SortAscending,
		"desc": db.SortDescending,
		"-1":   db.SortDescending,
	}

	if value, ok := sortMap[lowerStr]; ok {
		return value
	}
	return db.SortAscending
}

func CreateFilterFromQueryParam(d rules.Queryable, params url.Values) db.FindFilter {
	filter := db.FindFilter{}

	filter.Page = parseIntFromString(params.Get("page"))
	filter.PerPage = parseIntFromString(params.Get("perPage"))

	if searchStr := params.Get("q"); searchStr != "" {
		filter.Query = db.BuildSearchQueryFromString(d, searchStr)
	}

	filter.Sort.Field = params.Get("orderBy")
	filter.Sort.Order = getSortFromString(params.Get("sort"))

	return filter
}
