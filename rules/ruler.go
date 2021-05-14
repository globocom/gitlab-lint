// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules

import "github.com/xanzy/go-gitlab"

type Ruler interface {
	Run(client *gitlab.Client, p *gitlab.Project) bool
	GetSlug() string
	GetLevel() string
}

type RulerImpl struct {
	ID          string `json:"ruleId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       string `json:"level"`
}

func (e RulerImpl) GetSlug() string {
	return e.ID
}

func (e RulerImpl) GetLevel() string {
	return e.Level
}
