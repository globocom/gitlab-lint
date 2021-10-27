// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules

import "github.com/xanzy/go-gitlab"

type EmptyRepository struct {
	Description string `json:"description"`
	ID          string `json:"ruleId"`
	Level       string `json:"level"`
	Name        string `json:"name"`
}

func (e *EmptyRepository) Run(c *gitlab.Client, p *gitlab.Project) bool {
	return p.EmptyRepo
}

func (e *EmptyRepository) GetSlug() string {
	return "empty-repository"
}

func (e *EmptyRepository) GetLevel() string {
	return LevelError
}

func (e *EmptyRepository) GetName() string {
	return e.Name
}

func (e *EmptyRepository) GetDescription() string {
	return e.Description
}

func NewEmptyRepository() Ruler {
	e := &EmptyRepository{
		Name:        "Empty Repository",
		Description: "",
	}
	e.ID = e.GetSlug()
	e.Level = e.GetLevel()
	return e
}
