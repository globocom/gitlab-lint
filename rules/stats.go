// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules

import "time"

type Stats struct {
	CreatedAt            time.Time        `json:"createdAt" bson:"createdAt"`
	GitlabProjectsCount  int              `json:"gitlabProjectsCount" bson:"gitlabProjectsCount"`
	ProjectsCount        int              `json:"projectsCount" bson:"projectsCount"`
	RegisteredRulesCount int              `json:"registeredRulesCount" bson:"registeredRulesCount"`
	RulesCount           int              `json:"rulesCount" bson:"rulesCount"`
	LevelsCount          map[string]int32 `json:"levelsCount" bson:"levelsCount"`
}

func (s Stats) Cast() Queryable {
	return &s
}

func (s Stats) GetCollectionName() string {
	return "statistics"
}

func (s Stats) GetSearchableFields() []string {
	return nil
}
