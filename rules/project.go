// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules

import "github.com/xanzy/go-gitlab"

// FIXME
type Project struct {
	*gitlab.Project `json:",inline" bson:",inline"`
	Rules           map[string]int `json:"rules" bson:"rules"`
}

type Projects []Project

func (p Project) Cast() Queryable {
	return &p
}

func (p Project) GetCollectionName() string {
	return "projects"
}
